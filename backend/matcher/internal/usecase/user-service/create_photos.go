package user_service

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/google/uuid"
)

func (c *UseCase) createPhotos(ctx context.Context, userID uuid.UUID, requests []user.PhotoUploadRequest) ([]user.Photo, error) {
	photos := make([]user.Photo, 0, len(requests))
	for _, req := range requests {
		photoID, err := c.uuidProvider.GenerateV7()
		if err != nil {
			return nil, errors.WrapFail(err, "generate photo id")
		}

		photo := user.Photo{
			UserID:      userID,
			OrderNumber: req.OrderNumber,
			ObjectKey:   fmt.Sprintf("users/%s/slot/%d/%s.jpg", userID.String(), req.OrderNumber, photoID),
			MimeType:    req.MimeType,
		}

		uploadURL, err := c.photoURLGenerator.GenerateUpload(ctx, photo.ObjectKey, photo.MimeType)
		if err != nil {
			return nil, errors.Wrapf(
				err,
				"photoURLGenerator.GenerateUpload with %v and user with %v",
				errors.Token("order_number", photo.OrderNumber),
				errors.Token("user_id", photo.UserID),
			)
		}

		photo.UploadURL = &uploadURL

		photos = append(photos, photo)
	}

	return photos, nil
}
