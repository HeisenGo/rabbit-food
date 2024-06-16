package tcp

import (
	"fmt"
	"net"
	"rabbit-food/server/config"
	"rabbit-food/server/internal/server"
	"rabbit-food/server/internal/server/handlers"
	"rabbit-food/server/services"
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
	fmt.Println("Server started on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			//logger.Error("Error accepting connection:", err)
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go newServer.HandleConnection(conn)
	}
}
