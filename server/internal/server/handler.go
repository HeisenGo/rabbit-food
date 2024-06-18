package server

import (
	"context"
	"fmt"
	"net"
	"server/internal/server/handlers"
)

type Server struct {
	userHandler *handlers.UserHandler
}

func NewServer(userHandler *handlers.UserHandler) *Server {
	return &Server{userHandler}
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
		default:
			fmt.Println("default option!")
			conn.Write([]byte("incorrect option!"))
		}
	}
}
