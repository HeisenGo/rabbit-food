package tcp

import (
	"context"
	"fmt"
	"net"
	"server/config"
	"server/internal/server"
	"server/internal/server/handlers"
	"server/services"
)

func Run(cfg config.Server, app *services.AppContainer) {
	userHandler := handlers.NewUserHandler(*app.UserService)
	newServer := server.NewServer(userHandler)
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
