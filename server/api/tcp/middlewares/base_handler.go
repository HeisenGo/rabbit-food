package middleware

import (
	"context"
	"net"
	"server/internal/protocol/tcp"
)

type HandlerFunc func(context.Context, net.Conn, *tcp.Request)
