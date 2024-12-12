package main

import (
	"context"
	"log"
	"net/http"

	"github.com/brothergiez/golang-service-registry-pattern.git/internal/config"
	"github.com/brothergiez/golang-service-registry-pattern.git/internal/database"
	"github.com/brothergiez/golang-service-registry-pattern.git/internal/registry"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	client, err := database.ConnectMongo(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	db := client.Database(cfg.DatabaseName)
	repo := registry.NewMongoRepository(db, cfg.CollectionName)
	serviceRegistry := registry.NewServiceRegistry(repo)
	handler := registry.NewHandler(serviceRegistry)

	router := mux.NewRouter()
	registry.RegisterRoutes(router, handler)

	log.Printf("Starting server on port %s", cfg.ServerPort)

	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Failed to start server : %v", err)
	}
}
