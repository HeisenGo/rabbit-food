package handlers

import (
	"context"
	"net"
	"server/internal/protocol/tcp"
)

type Handler interface {
	ServeTCP(ctx context.Context, conn net.Conn, TCPReq *tcp.Request)
}

type HandlerFunc func(context.Context, net.Conn, *tcp.Request)

func (f HandlerFunc) ServeTCP(ctx context.Context, conn net.Conn, TCPReq *tcp.Request) {
	f(ctx, conn, TCPReq)
}
