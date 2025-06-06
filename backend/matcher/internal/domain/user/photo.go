//go:generate mockgen -source=photo.go -destination=../../mocks/user/photo_mock.go -typed

package user

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

const (
	uploadURLLiveTime   = 15 * time.Minute
	downloadURLLiveTime = 30 * time.Minute
)

var (
	ErrPhotoNotFound  = errors.Error("photo not found")
	ErrPhotosNotFound = errors.Error("photos not found")
)

func NewObjectStoragePhotoURLGenerator(client *minio.Client, bucket string) *objectStoragePhotoURLGenerator {
	return &objectStoragePhotoURLGenerator{
		client: client,
		bucket: bucket,
	}
}

type objectStoragePhotoURLGenerator struct {
	client *minio.Client
	bucket string
}

func (g *objectStoragePhotoURLGenerator) GenerateUpload(ctx context.Context, key string, _ string) (string, error) {
	expires := uploadURLLiveTime
	uploadURL, err := g.client.PresignedPutObject(ctx, g.bucket, key, expires)
	if err != nil {
		return "", errors.WrapFailf(
			err,
			"generate signed upload url %v %v",
			errors.Token("bucket", g.bucket),
			errors.Token("key", key),
		)
	}

	return uploadURL.String(), nil
}

func (g *objectStoragePhotoURLGenerator) GenerateDownload(_ context.Context, key string) (string, error) {
	u := &url.URL{
		Scheme: "https",
		Host:   "photos.testing.couply.ru",
		Path:   fmt.Sprintf("/%s/%s", g.bucket, key),
	}
	q := u.Query()
	q.Set("response-cache-control", "public, max-age=604800, immutable")
	u.RawQuery = q.Encode()
	return u.String(), nil
}

type PhotoURLGenerator interface {
	GenerateUpload(ctx context.Context, key string, contentType string) (string, error)
	GenerateDownload(ctx context.Context, key string) (string, error)
}

type PhotoUploadRequest struct {
	OrderNumber int32
	MimeType    string
}

type Photo struct {
	UserID      uuid.UUID  `db:"user_id"`
	OrderNumber int32      `db:"order_number"`
	ObjectKey   string     `db:"object_key"`
	MimeType    string     `db:"mime_type"`
	UploadedAt  *time.Time `db:"uploaded_at"`

	UploadURL   *string
	DownloadURL *string
}

func (x *Photo) GetDownloadURL(ctx context.Context, gen PhotoURLGenerator) error {
	if x == nil || x.UploadedAt == nil {
		return nil
	}

	downloadURL, err := gen.GenerateDownload(ctx, x.ObjectKey)
	if err != nil {
		return errors.WrapFailf(
			err,
			"generate download url for photo with %v and user with %v",
			errors.Token("order_number", x.OrderNumber),
			errors.Token("user_id", x.UserID),
		)
	}

	x.DownloadURL = &downloadURL

	return nil
}
