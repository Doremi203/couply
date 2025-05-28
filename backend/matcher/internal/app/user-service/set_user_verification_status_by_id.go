package user_service

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) SetUserVerificationStatusByIDV1(
	ctx context.Context,
	req *desc.SetUserVerificationStatusByIDV1Request,
) (*desc.SetUserVerificationStatusByIDV1Response, error) {
	var verificationStatus user.VerificationStatus
	switch req.GetStatus() {
	case desc.VerificationStatus_PASS:
		verificationStatus = user.VerificationStatusPass
	case desc.VerificationStatus_FAIL:
		verificationStatus = user.VerificationStatusFail
	case desc.VerificationStatus_MANUAL:
		verificationStatus = user.VerificationStatusManual
	default:
		verificationStatus = user.VerificationStatusUnknown
	}

	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		i.logger.Error(errors.WrapFail(err, "parse user id"))
	}

	err = i.usecase.SetUserVerificationStatusByID(ctx, userID, verificationStatus)
	switch {
	case errors.Is(err, user.ErrUserNotFound):
		return nil, status.Error(codes.NotFound, user.ErrUserNotFound.Error())
	case err != nil:
		return nil, err
	}

	return &desc.SetUserVerificationStatusByIDV1Response{
		Success: true,
	}, nil
}
