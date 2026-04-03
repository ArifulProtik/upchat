package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"ArifulProtik/UpChat/internal"
	"ArifulProtik/UpChat/internal/controller"
	"ArifulProtik/UpChat/internal/ent"
	"ArifulProtik/UpChat/internal/ent/migrate"
	"ArifulProtik/UpChat/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	config, err := internal.LoadConfig()
	if err != nil {
		log.Fatalf("[main] Failed to load config: %v", err)
	}
	dbClient, err := ent.Open(config.DatabaseDriver, config.DatabaseURL)
	if err != nil {
		log.Fatalf("[main] Failed to open database: %v", err)
	}
	err = dbClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("[main] Failed to create schema: %v", err)
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
