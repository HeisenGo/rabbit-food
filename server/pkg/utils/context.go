package utils

import (
	"context"
	"errors"
	"server"
)

func GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value(server.UserIDKey).(uint)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}
	return userID, nil
}
