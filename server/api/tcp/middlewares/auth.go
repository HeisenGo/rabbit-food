package middleware

import (
	"context"
	"net"
	"server"
	"server/api/tcp/handlers"
	"server/internal/protocol/tcp"
	"server/pkg/jwt"
	"strings"
)

func AuthMiddleware(secret string, next handlers.Handler) handlers.Handler {
	return handlers.HandlerFunc(func(ctx context.Context, conn net.Conn, TCPReq *tcp.Request) {
		token := strings.TrimPrefix(TCPReq.Header["Authorization"], "Bearer ")
		if token == "" {
			tcp.Error(conn, tcp.StatusUnauthorized, nil, "missing auth token")
			return
		}
		userID, err := jwt.ParseToken(token, []byte(secret))
		if err != nil {
			tcp.Error(conn, tcp.StatusUnauthorized, nil, "invalid auth token")
			return
		}

		ctx = context.WithValue(ctx, server.UserIDKey, userID)
		next.ServeTCP(ctx, conn, TCPReq)
	})
}
