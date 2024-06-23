package tcp

import (
	"context"
	"net"
	"server/api/tcp/handlers"
	"server/config"
	"server/internal/server"
	"server/services"
	"server/pkg/logger"
)

func Run(cfg config.Server, app *services.AppContainer,log *logger.CustomLogger) {
	
	userHandler := handlers.NewUserHandler(*app.UserService,log)
	newServer := server.NewServer(userHandler,log) 
	listener, err := net.Listen("tcp", ":"+cfg.Port)
	log.Info("listening!")
	if err != nil {
		log.Error("Error starting Server")
		return
	}
	defer listener.Close()
	log.Info("Server started on port 8080")
	ctx := context.Background()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Warn("Error accepting connection:", err)
			continue
		}
		go newServer.HandleConnection(ctx, conn)
	}
}
