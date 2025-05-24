package postgres

import (
	"bytes"
	"github.com/google/uuid"
)

const (
	likesTableName   = "likes"
	matchesTableName = "matches"

	senderIDColumnName   = "sender_id"
	receiverIDColumnName = "receiver_id"
	messageColumnName    = "message"
	statusColumnName     = "status"
	createdAtColumnName  = "created_at"

	firstUserIDColumnName  = "first_user_id"
	secondUserIDColumnName = "second_user_id"
)

var (
	likesColumns   = []string{senderIDColumnName, receiverIDColumnName, messageColumnName, statusColumnName, createdAtColumnName}
	matchesColumns = []string{firstUserIDColumnName, secondUserIDColumnName, createdAtColumnName}
)

// orderUserIDs guarantees the same order for all likes
func orderUserIDs(id1, id2 uuid.UUID) (uuid.UUID, uuid.UUID) {
	if bytes.Compare(id1[:], id2[:]) < 0 {
		return id1, id2
	}
	return id2, id1
}
