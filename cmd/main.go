package main

import (
	"log"
	"os"
	"quasar/config"
	"quasar/internal/handlers"
	"quasar/internal/repository"
	"quasar/internal/services"

	"github.com/joho/godotenv"
)

func main() {
	// Load env variables if not production
	var port string
	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			return
		}

	}
	port = ":" + os.Getenv("SERVER_PORT")

	// Initialize the server, inyection dependecy
	service := services.NewServiceSatelliteImpl(repository.NewListRepositoryImpl(), services.NewQuasarImpl())
	handler := handlers.NewHandlerImpl(service)
	var server = config.NewServer(handler, port)

	// Run the server
	server.Run()
}
