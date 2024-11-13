package main

import (
	"log"

	"github.com/mirjalilova/websiteOfEverest/config"
	"github.com/mirjalilova/websiteOfEverest/internal/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	app.Run(cfg)

}
