package auth

import (
	"context"
	"net/http"
	"server"
)

func SetUserIDInContext(r *http.Request, userID uint) *http.Request {
	ctx := context.WithValue(r.Context(), server.UserIDKey, userID)
	return r.WithContext(ctx)
}
