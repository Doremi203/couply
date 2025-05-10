package user_service

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	userpkg "github.com/Doremi203/couply/backend/auth/pkg/user"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

func (c *UseCase) createPhotos(ctx context.Context, userID userpkg.ID, requests []user.PhotoUploadRequest) ([]user.Photo, error) {
	photos := make([]user.Photo, 0, len(requests))
	for _, req := range requests {
		photo := user.Photo{
			UserID:      userID,
			OrderNumber: req.OrderNumber,
			ObjectKey:   fmt.Sprintf("users/%s/slots/%d.%s", userID, req.OrderNumber, req.MimeType),
			MimeType:    req.MimeType,
		}

		uploadURL, err := c.photoURLGenerator.GenerateUpload(ctx, photo.ObjectKey, photo.MimeType)
		if err != nil {
			return nil, errors.WrapFailf(
				err,
				"generate upload url for photo with %v and user with %v",
				errors.Token("order_number", photo.OrderNumber),
				errors.Token("user_id", photo.UserID),
			)
		}

		photo.UploadURL = &uploadURL

		photos = append(photos, photo)
	}

	return photos, nil
}

func (c *UseCase) downloadablePhotos(ctx context.Context, u user.User) ([]user.Photo, error) {
	ret := make([]user.Photo, 0, len(u.Photos))
	for _, photo := range u.Photos {
		if photo.UploadedAt == nil {
			continue
		}
		downloadURL, err := c.photoURLGenerator.GenerateDownload(ctx, photo.ObjectKey)
		if err != nil {
			return nil, errors.WrapFailf(
				err,
				"generate download url for photo with %v and user with %v",
				errors.Token("order_number", photo.OrderNumber),
				errors.Token("user_id", photo.UserID),
			)
		}

		photo.DownloadURL = &downloadURL
		ret = append(ret, photo)
	}

	return ret, nil
}
