package server

import (
	"context"
	"fmt"
	"net"
	"server/api/tcp/handlers"
	"server/internal/protocol/tcp"
	"strings"
)

type Server struct {
	authHandler *handlers.AuthHandler
	userHandler *handlers.UserHandler
}

func NewServer(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler) *Server {
	return &Server{
		authHandler: authHandler,
		userHandler: userHandler,
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
			return
		}
		buffer = buffer[:n]

		requestData, err := tcp.DecodeTCPRequest(buffer)
		if err != nil {
			//logger.Error("Error decoding register request:", err)
			fmt.Println("request format is not correct.", err)
			return
		}
		allRoutes := strings.Split(requestData.Location, "/")
		route := allRoutes[0]
		requestData.Location = strings.Join(allRoutes[1:], "/")
		switch route {
		case "auth":
			s.authHandler.AuthRouter(ctx, conn, requestData)
		case "user":
			s.userHandler.UserRouter(ctx, conn, requestData)
		default:
			fmt.Println("default option!")
			conn.Write([]byte("incorrect option!"))
		}
	}
}
