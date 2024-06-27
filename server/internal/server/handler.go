package server

import (
	"context"
	"fmt"
	"net"
	"server/api/tcp/handlers"
	"server/internal/protocol/tcp"
	"server/pkg/utils"
)

type Server struct {
	authHandler        *handlers.AuthHandler
	walletHandler      *handlers.WalletHandler
	restaurantHandler  *handlers.RestaurantHandler
	userProfileHandler *handlers.UserProfileHandler // Add this line
}

func NewServer(authHandler *handlers.AuthHandler, walletHandler *handlers.WalletHandler, restaurantHandler *handlers.RestaurantHandler, userProfileHandler *handlers.UserProfileHandler) *Server { // Update this line
	return &Server{authHandler, walletHandler, restaurantHandler, userProfileHandler}
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
		fmt.Println(string(buffer))

		requestData, err := tcp.DecodeTCPRequest(buffer)
		if err != nil {
			//logger.Error("Error decoding register request:", err)
			fmt.Println("request format is not correct.", err)
			return
		}
		firstRoute, postRoutes := utils.RouteSplitter(requestData.Location)
		requestData.Location = postRoutes
		switch firstRoute {
		case "auth":
			s.authHandler.ServeTCP(ctx, conn, requestData)
		case "wallets":
			s.walletHandler.ServeTCP(ctx, conn, requestData)
		case "restaurant":
			s.restaurantHandler.ServeTCP(ctx, conn, requestData)
		case "profile": // Add this case
			s.userProfileHandler.ServeTCP(ctx, conn, requestData)
		default:
			fmt.Println("default option!")
			conn.Write([]byte("incorrect option!"))
		}
	}
}
