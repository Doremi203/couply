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
		photo := user.Photo{
			UserID:      userID,
			OrderNumber: req.OrderNumber,
			ObjectKey:   fmt.Sprintf("users/%s/slot/%d.jpg", userID, req.OrderNumber),
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
