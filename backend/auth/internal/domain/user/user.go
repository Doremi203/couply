package user

import (
	"regexp"
	"time"

	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	uuidgen "github.com/Doremi203/couply/backend/auth/pkg/uuid"
	"github.com/google/uuid"
)

type ID uuid.UUID

func (id ID) String() string {
	return uuid.UUID(id).String()
}

type Email string

func NewPhone(v string) (Phone, error) {
	re := regexp.MustCompile(`^\+\d{10,15}$`)
	if !re.MatchString(v) {
		return "", errors.Error("invalid phone format, must be in E.164 format")
	}

	return Phone(v), nil
}

type Phone string

type CreatedAt time.Time

func New(
	uuidProvider uuidgen.Provider,
	email Email,
	password pswrd.HashedPassword,
) (User, error) {
	id, err := uuidProvider.GenerateV7()
	if err != nil {
		return User{}, errors.WrapFail(err, "generate user id")
	}

	return User{
		ID:       ID(id),
		Email:    email,
		Password: password,
	}, nil
}

type User struct {
	ID       ID
	Email    Email
	Phone    Phone
	Password pswrd.HashedPassword

	OAuthAccounts []OAuthAccount
}
