package user

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/google/uuid"
)

type UID uuid.UUID

type UIDGenerator interface {
	Generate() (UID, error)
}

type UUIDV7BasedUIDGenerator struct {
}

func (g *UUIDV7BasedUIDGenerator) Generate() (UID, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return UID{}, errors.Wrap(err, "failed to generate uuid v7")
	}
	return UID(uuid), nil
}
