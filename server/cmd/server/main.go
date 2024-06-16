package main

import (
	"flag"
	"log"
	"server/api/tcp"
	"server/config"
	"server/services"
)

var envFilePath = flag.String("envpath", "", "configuration path")

func main() {
	cfg := readConfig()

	app, err := services.NewAppContainer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	tcp.Run(cfg.Server, app)

}

func readConfig() config.Config {
	flag.Parse()
	return config.ReadConfig(*envFilePath)
}
