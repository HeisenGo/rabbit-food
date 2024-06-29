package tcp

import (
	"context"
	"fmt"
	"net"
	"server/api/tcp/handlers"
	"server/config"
	"server/internal/server"
	"server/services"
)

func Run(cfg config.Server, app *services.AppContainer) {
	authHandler := handlers.NewAuthHandler(*app.AuthService)
	walletHandler := handlers.NewWalletHandler(*app.WalletService)

	addressHandler := handlers.NewAddressHandler(*app.AddressService)
	restaurantHandler := handlers.NewRestaurantHandler(*app.RestaurantService, *app.UserService)
	newServer := server.NewServer(authHandler, walletHandler, restaurantHandler, addressHandler)
	listener, err := net.Listen("tcp", ":"+cfg.Port)
	fmt.Println("listening!")
	if err != nil {
		//logger.Error("Error starting server:", err)
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	//logger.Info("Server started on port 8080")
	fmt.Printf("Server started on port %v\n", cfg.Port)

	ctx := context.Background()
	for {
		conn, err := listener.Accept()
		if err != nil {
			//logger.Error("Error accepting connection:", err)
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go newServer.HandleConnection(ctx, conn)
	}
}
