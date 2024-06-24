package middleware

import (
	"context"
	"net"
	"server"
	"server/config"
	"server/internal/protocol/tcp"
	"server/pkg/jwt"
	"strings"
)

func AuthMiddleware(next HandlerFunc) HandlerFunc {
	return func(ctx context.Context, conn net.Conn, TCPReq *tcp.Request) {
		tokenString := strings.TrimPrefix(TCPReq.Header["Authorization"], "Bearer ")
		if tokenString == "" {
			tcp.Error(conn, tcp.StatusUnauthorized, nil, "missing auth token")
			return
		}
		cfg := config.Get()
		token, err := jwt.ParseToken(tokenString, []byte(cfg.Server.TokenSecret))
		if err != nil {
			tcp.Error(conn, tcp.StatusUnauthorized, nil, "invalid auth token")
			return
		}
		ctx = context.WithValue(ctx, server.UserIDKey, token.UserID)
		next(ctx, conn, TCPReq)
	}
}
