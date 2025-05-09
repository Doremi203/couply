package phoneconfirm

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/internal/domain/hash"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/timeprovider"
	"github.com/google/uuid"
)

func NewRequest(
	codeGenerator CodeGenerator,
	hashProvider hash.Provider,
	userID user.ID,
	phone user.Phone,
) (Request, error) {
	code, err := codeGenerator.Generate()
	if err != nil {
		return Request{}, errors.WrapFail(err, "generate secure code")
	}

	hashedCode, err := hashProvider.Hash(code.value)
	if err != nil {
		return Request{}, errors.WrapFail(err, "hash code")
	}

	return Request{
		UserID:     userID,
		Phone:      phone,
		Code:       code,
		HashedCode: hashedCode,
	}, nil
}

type Request struct {
	UserID     user.ID
	Phone      user.Phone
	Code       Code
	HashedCode HashedCode
	CreatedAt  time.Time
}

func (r *Request) Expired(provider timeprovider.Provider) bool {
	return provider.Now().After(r.CreatedAt.Add(r.Code.ExpiresIn))
}

type ID uuid.UUID
