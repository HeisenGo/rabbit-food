package server

import (
	"context"
	"net"
	"server/internal/server/handlers"
	"server/pkg/logger"
)

type Server struct {
	userHandler *handlers.UserHandler
	logger *logger.CustomLogger
}

func NewServer(userHandler *handlers.UserHandler,logger *logger.CustomLogger) *Server {
	return &Server{userHandler:userHandler,
					logger: logger,
				}
}

func (s *Server) HandleConnection(ctx context.Context, conn net.Conn) {
	ctx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()
	defer conn.Close()

	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			s.logger.Debug("Nothing To Read ",err)
			return
		}

		// Assume the first byte indicates the type of request
		switch buffer[0] {
		case '1': // Register request
			s.userHandler.HandleRegister(ctx, conn, buffer[1:n])
			s.logger.Info("Register Handled !")
		// Add other cases for different requests
		default:
			s.logger.Warn(" Wrong Request")
			conn.Write([]byte("incorrect option!"))
		}
	}
}
