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
	authHandler       *handlers.AuthHandler
	walletHandler     *handlers.WalletHandler
	restaurantHandler *handlers.RestaurantHandler
	addressHandler    *handlers.AddressHandler
}

func NewServer(authHandler *handlers.AuthHandler, walletHandler *handlers.WalletHandler, restaurantHandler *handlers.RestaurantHandler, addressHandler *handlers.AddressHandler, userHandler *handlers.UserHandler) *Server {
	return &Server{authHandler, walletHandler, restaurantHandler, addressHandler}

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
		case "users":
			s.addressHandler.ServeTCP(ctx, conn, requestData)
		case "restaurants":
			s.restaurantHandler.ServeTCP(ctx, conn, requestData)
		default:
			fmt.Println("default option!")
			conn.Write([]byte("incorrect option!"))
		}
	}
}
