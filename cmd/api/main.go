package main

import (
	"ArifulProtik/UpChat/internal"
	"ArifulProtik/UpChat/internal/controller"
	"ArifulProtik/UpChat/internal/service"
	"log"
	"log/slog"
	"os"
)

func main() {
	config, err := internal.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	logHandler := slog.HandlerOptions{
		AddSource: true,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &logHandler))
	newService := service.New(logger)
	newController := controller.New(logger, newService)
	server := internal.New(config)
	internal.SetupRoutes(server.Route, newController)
	server.Run()
}
