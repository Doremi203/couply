package blocker_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"

	desc "github.com/Doremi203/couply/backend/blocker/gen/api/blocker-service/v1"
	"github.com/Doremi203/couply/backend/blocker/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ReportUserV1(ctx context.Context, in *desc.ReportUserV1Request) (*desc.ReportUserV1Response, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req, err := dto.PBToReportUserRequest(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response, err := i.usecase.ReportUser(ctx, req)
	switch {
	case errors.Is(err, blocker.ErrDuplicateUserBlock):
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	case err != nil:
		return nil, err
	}

	return dto.ReportUserResponseToPB(response), nil
}
