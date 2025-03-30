package domain

import (
	"time"

	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Photo struct {
	UserID      int64     `db:"user_id"`
	OrderNumber int32     `db:"order_number"`
	URL         string    `db:"url"`
	MimeType    string    `db:"mime_type"`
	UploadedAt  time.Time `db:"uploaded_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func PhotoToPB(photo *Photo) *desc.Photo {
	return &desc.Photo{
		UserId:      photo.UserID,
		OrderNumber: photo.OrderNumber,
		Url:         photo.URL,
		MimeType:    photo.MimeType,
		UploadedAt:  timestamppb.New(photo.UploadedAt),
		UpdatedAt:   timestamppb.New(photo.UpdatedAt),
	}
}

func PBToPhoto(photo *desc.Photo) *Photo {
	return &Photo{
		UserID:      photo.UserId,
		OrderNumber: photo.OrderNumber,
		URL:         photo.Url,
		MimeType:    photo.MimeType,
		UploadedAt:  photo.UploadedAt.AsTime(),
		UpdatedAt:   photo.UpdatedAt.AsTime(),
	}
}

func PhotoSliceToPB(photos []*Photo) []*desc.Photo {
	photosPB := make([]*desc.Photo, 0, len(photos))

	for _, photo := range photos {
		photosPB = append(photosPB, PhotoToPB(photo))
	}

	return photosPB
}

func PBToPhotoSlice(photos []*desc.Photo) []*Photo {
	photosDomain := make([]*Photo, 0, len(photos))

	for _, photo := range photos {
		photosDomain = append(photosDomain, PBToPhoto(photo))
	}

	return photosDomain
}
