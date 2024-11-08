package main

import (
	"github.com/YurcheuskiRadzivon/HSC-pattern/api/routes"
	"github.com/YurcheuskiRadzivon/HSC-pattern/config"
	"github.com/YurcheuskiRadzivon/HSC-pattern/internal/hsc_helpers"
	"log"
	"path/filepath"
)

func main() {
	cfgPath := filepath.Join("..", "..", "config", "config.yaml")
	dsnStr, err := config.GetConfig(cfgPath)
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	userHandler, err := hsc_helpers.InitializeComponentsUser(dsnStr)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	propertyHandler, err := hsc_helpers.InitializeComponentsProperty(dsnStr)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	app := routes.NewFiberRouter(userHandler, propertyHandler)
	app.Listen(":8080")
}
