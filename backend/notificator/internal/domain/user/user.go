package user

import "github.com/google/uuid"

type ID uuid.UUID

func (id ID) String() string {
	return uuid.UUID(id).String()
}
