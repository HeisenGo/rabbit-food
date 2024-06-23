package main

import (
	"flag"
	"log"
	"server/api/tcp"
	"server/config"
	"server/services"
	"server/pkg/logger"
		
)

var envFilePath = flag.String("envpath", "", "configuration path")

func main() {
	config := &logger.LoggerConfig{
		WriteToConsole: true,
		WriteToFile:    true,
		FilePath:       "server.log",
	}
	customlog, err := logger.NewCustomLogger(config)
	if err != nil {
		log.Fatalf("Could not initialize logger: %v", err)
	}
	cfg := readConfig()
	app, err := services.NewAppContainer(cfg,customlog)
	if err!=nil{
		customlog.Error("Error In Running App Container.")
	}
	tcp.Run(cfg.Server, app,customlog)

}

func readConfig() config.Config {
	flag.Parse()
	return config.ReadConfig(*envFilePath)
}
