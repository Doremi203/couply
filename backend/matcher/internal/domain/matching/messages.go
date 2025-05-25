package matching

import (
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

type LikeMessage struct {
	ReceiverID uuid.UUID `json:"receiver_id"`
	Message    string    `json:"message"`
}

type MatchMessage struct {
	FirstUserID  uuid.UUID `json:"first_user_id"`
	SecondUserID uuid.UUID `json:"second_user_id"`
}

func (m *LikeMessage) String() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (m *MatchMessage) String() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(bytes)
}
