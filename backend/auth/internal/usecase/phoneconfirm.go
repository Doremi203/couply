package usecase

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/internal/domain/hash"
	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/sms"
	"github.com/Doremi203/couply/backend/auth/pkg/timeprovider"
)

func NewPhoneConfirmation(
	smsSender sms.Sender,
	codeGenerator phoneconfirm.CodeGenerator,
	hashProvider hash.Provider,
	confirmRepo phoneconfirm.Repo,
) PhoneConfirmation {
	return PhoneConfirmation{
		smsSender:     smsSender,
		codeGenerator: codeGenerator,
		hashProvider:  hashProvider,
		confirmRepo:   confirmRepo,
		timeProvider:  timeprovider.ProviderFunc(time.Now),
	}
}

type PhoneConfirmation struct {
	smsSender     sms.Sender
	codeGenerator phoneconfirm.CodeGenerator
	hashProvider  hash.Provider
	timeProvider  timeprovider.Provider
	confirmRepo   phoneconfirm.Repo
}

var ErrPendingConfirmationRequestAlreadyExists = errors.Error("phone confirmation request already exists")

func (u PhoneConfirmation) SendCodeV1(ctx context.Context, userID user.ID, phone user.Phone) (phoneconfirm.Request, error) {
	existingReq, err := u.confirmRepo.GetRequest(ctx, userID, phone)
	if err != nil {
		return phoneconfirm.Request{}, errors.WrapFailf(
			err,
			"get existing request for %v",
			errors.Token("user_id", userID),
		)
	}
	if existingReq != nil && !existingReq.Expired(u.timeProvider) {
		return phoneconfirm.Request{}, ErrPendingConfirmationRequestAlreadyExists
	}

	request, err := phoneconfirm.NewRequest(
		u.codeGenerator,
		u.hashProvider,
		userID,
		phone,
	)
	if err != nil {
		return phoneconfirm.Request{}, errors.WrapFail(err, "create phone confirmation request")
	}

	err = u.confirmRepo.UpsertRequest(ctx, request)
	if err != nil {
		return phoneconfirm.Request{}, errors.WrapFail(err, "upsert phone confirmation request")
	}

	err = u.smsSender.Send(ctx, request.Code.Value(), string(phone))
	if err != nil {
		return phoneconfirm.Request{}, errors.WrapFail(err, "send phone code")
	}

	return request, nil
}

func (u PhoneConfirmation) ConfirmV1(ctx context.Context, confirmID phoneconfirm.ID) {}
