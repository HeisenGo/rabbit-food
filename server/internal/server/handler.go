package server

import (
	"context"
	"fmt"
	"net"
	"server/api/tcp/handlers"
)

type Server struct {
	userHandler *handlers.UserHandler
	authHandler *handlers.AuthHandler
}

func NewServer(userHandler *handlers.UserHandler,authHandler *handlers.AuthHandler) *Server {
	return &Server{userHandler,authHandler}
}

func (s *Server) HandleConnection(ctx context.Context, conn net.Conn) {
	ctx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()
	defer conn.Close()

	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}

		// Assume the first byte indicates the type of request
		switch buffer[0] {
		case '1': // Register request
			s.userHandler.HandleRegister(ctx, conn, buffer[1:n])
		// Add other cases for different requests
		case '2':
			s.authHandler.HandleLogin(ctx,conn,buffer[1:n])
		default:
			fmt.Println("default option!")
			conn.Write([]byte("incorrect option!"))
		}
	}
}
