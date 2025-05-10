package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ConfirmPhotosUploadV1(ctx context.Context, in *desc.ConfirmPhotosUploadV1Request) (*desc.ConfirmPhotosUploadV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := i.usecase.ConfirmPhotosUpload(ctx, in.GetOrderNumbers())
	if err != nil {
		i.logger.Error(errors.Wrap(err, "confirm photos upload v1 failed"))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &desc.ConfirmPhotosUploadV1Response{}, nil
}
