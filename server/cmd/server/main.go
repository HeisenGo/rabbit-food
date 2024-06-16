package main

import (
	"flag"
	"log"
	"rabbit-food/server/api/tcp"
	"rabbit-food/server/config"
	"rabbit-food/server/services"
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
