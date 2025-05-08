package matching

import (
	"bytes"

	"github.com/google/uuid"
)

// orderUserIDs гарантирует один порядок айди пользователей
func orderUserIDs(id1, id2 uuid.UUID) (uuid.UUID, uuid.UUID) {
	if bytes.Compare(id1[:], id2[:]) < 0 {
		return id1, id2
	}
	return id2, id1
}
