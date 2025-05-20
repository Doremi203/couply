package usecase

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/internal/domain/hash"
	"github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/auth/pkg/sms"
	"github.com/Doremi203/couply/backend/auth/pkg/timeprovider"
	"github.com/Doremi203/couply/backend/auth/pkg/tx"
)

func NewPhoneConfirmation(
	smsSender sms.Sender,
	codeGenerator phoneconfirm.CodeGenerator,
	hashProvider hash.Provider,
	confirmRepo phoneconfirm.Repo,
	userRepo user.Repo,
	txProvider tx.Provider,
) PhoneConfirmation {
	return PhoneConfirmation{
		smsSender:     smsSender,
		codeGenerator: codeGenerator,
		hashProvider:  hashProvider,
		confirmRepo:   confirmRepo,
		userRepo:      userRepo,
		txProvider:    txProvider,
		timeProvider:  timeprovider.ProviderFunc(time.Now),
	}
}

type PhoneConfirmation struct {
	smsSender     sms.Sender
	codeGenerator phoneconfirm.CodeGenerator
	hashProvider  hash.Provider
	timeProvider  timeprovider.Provider
	confirmRepo   phoneconfirm.Repo
	userRepo      user.Repo
	txProvider    tx.Provider
}

var ErrPendingConfirmationRequestAlreadyExists = errors.Error("phone confirmation request already exists")
var ErrPhoneAlreadyConfirmed = errors.Error("phone already confirmed for some user")
var ErrUnsupportedPhoneOperator = errors.Error("phone operator not supported")

func (u PhoneConfirmation) SendCodeV1(ctx context.Context, userID user.ID, phone user.Phone) (phoneconfirm.Request, error) {
	_, err := u.userRepo.GetByAny(ctx, user.GetByAnyParams{Phone: phone})
	switch {
	case err == nil:
		return phoneconfirm.Request{}, ErrPhoneAlreadyConfirmed

	case errors.Is(err, user.ErrNotFound):
	// continue

	case err != nil:
		return phoneconfirm.Request{}, errors.WrapFail(err, "get user by phone")
	}

	existingReq, err := u.confirmRepo.GetRequest(ctx, userID, phone)
	if err != nil && !errors.Is(err, phoneconfirm.ErrRequestNotFound) {
		return phoneconfirm.Request{}, errors.WrapFailf(
			err,
			"get existing request for %v",
			errors.Token("user_id", userID),
		)
	}
	if !existingReq.Expired(u.timeProvider) {
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

	err = u.smsSender.Send(ctx, string(request.Code.Value()), string(phone))
	switch {
	case errors.Is(err, sms.ErrUnsupportedPhoneOperator):
		return phoneconfirm.Request{}, ErrUnsupportedPhoneOperator
	case err != nil:
		return phoneconfirm.Request{}, errors.WrapFail(err, "send phone code")
	}

	return request, nil
}

var ErrNoPendingConfirmRequestExists = errors.Error(" no pending phone confirmation request exists")
var ErrConfirmCodeExpired = errors.Error("phone confirmation code expired")
var ErrIncorrectConfirmationCode = errors.Error("incorrect confirmation code")

func (u PhoneConfirmation) ConfirmV1(
	ctx context.Context,
	logger log.Logger,
	userID user.ID,
	phone user.Phone,
	code phoneconfirm.CodeValue,
) error {
	request, err := u.confirmRepo.GetRequest(ctx, userID, phone)
	switch {
	case errors.Is(err, phoneconfirm.ErrRequestNotFound):
		return ErrNoPendingConfirmRequestExists
	case err != nil:
		return errors.WrapFailf(
			err,
			"get phone confirmation request for %v",
			errors.Token("user_id", userID),
		)
	}
	if request.Expired(u.timeProvider) {
		return ErrConfirmCodeExpired
	}

	err = u.hashProvider.Verify(string(code), request.HashedCode)
	switch {
	case errors.Is(err, hash.ErrNoMatch):
		return ErrIncorrectConfirmationCode
	case err != nil:
		return errors.WrapFail(err, "verify phone confirmation code")
	}

	ctx, err = u.txProvider.ContextWithTx(ctx, tx.IsolationSerializable)
	if err != nil {
		return errors.WrapFail(err, "create tx context")
	}
	defer func() {
		if err != nil {
			err := u.txProvider.RollbackTx(ctx)
			if err != nil {
				logger.Error(errors.WrapFail(err, "rollback confirm code tx"))
			}
		}
	}()

	_, err = u.userRepo.GetByAny(ctx, user.GetByAnyParams{Phone: phone})
	switch {
	case err == nil:
		return ErrPhoneAlreadyConfirmed

	case errors.Is(err, user.ErrNotFound):
	// continue

	case err != nil:
		return errors.WrapFail(err, "get user by phone")
	}

	err = u.userRepo.UpdatePhone(ctx, userID, phone)
	if err != nil {
		return errors.WrapFailf(err, "update phone for user with %v", errors.Token("id", userID))
	}

	err = u.confirmRepo.DeleteRequest(ctx, userID, phone)
	if err != nil {
		return errors.WrapFailf(err, "delete phone confirmation request for %v", errors.Token("id", userID))
	}

	err = u.txProvider.CommitTx(ctx)
	if err != nil {
		return errors.WrapFail(err, "commit confirm code tx")
	}

	return nil
}
