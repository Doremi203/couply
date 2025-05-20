package utils

import (
	"context"
	"fmt"

	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/google/uuid"
)

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userToken, ok := token.FromContext(ctx)
	if !ok {
		return uuid.Nil, fmt.Errorf("user token not found in context")
	}
	return userToken.GetUserID(), nil
}
