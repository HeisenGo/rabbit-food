package main

import (
	"flag"
	"log"
	"server/api/tcp"
	"server/config"
	"server/services"
)

func main() {
	cfg := readConfig()
	config.Set(cfg)

	app, err := services.NewAppContainer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	tcp.Run(cfg.Server, app)

}

var envFilePath = flag.String("envpath", "", ".env file path")

func readConfig() config.Config {
	flag.Parse()
	return config.ReadConfig(*envFilePath)
}
