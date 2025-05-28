package user_service

import (
	"context"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ConfirmPhotosUploadV1(ctx context.Context, in *desc.ConfirmPhotosUploadV1Request) (*desc.ConfirmPhotosUploadV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &desc.ConfirmPhotosUploadV1Response{}, nil

	err := i.usecase.ConfirmPhotosUpload(ctx, in.GetOrderNumbers())
	if err != nil {
		return nil, err
	}

	return &desc.ConfirmPhotosUploadV1Response{}, nil
}
