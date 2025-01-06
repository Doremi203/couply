package domain

import (
	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

type Photo struct {
	ID         int64     `db:"id"`
	UserID     int64     `db:"user_id"`
	URL        string    `db:"url"`
	MimeType   string    `db:"mime_type"`
	UploadedAt time.Time `db:"uploaded_at"`
}

func PhotoToPB(photo *Photo) *desc.Photo {
	return &desc.Photo{
		Id:         strconv.FormatInt(photo.ID, 10),
		Url:        photo.URL,
		MimeType:   photo.MimeType,
		UploadedAt: timestamppb.New(photo.UploadedAt),
	}
}

func PBToPhoto(photo *desc.Photo) *Photo {
	id, _ := strconv.Atoi(photo.Id)

	return &Photo{
		ID:         int64(id),
		URL:        photo.Url,
		MimeType:   photo.MimeType,
		UploadedAt: photo.UploadedAt.AsTime(),
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
