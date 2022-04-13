package config

import (
	"quasar/internal/contracts"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	engine   *fiber.App
	handlers contracts.IHandlerFiber
	Port     string
}

func NewServer(handlers contracts.IHandlerFiber, port string) *Server {
	// Initialize the server
	server := Server{
		engine:   fiber.New(),
		handlers: handlers,
		Port:     port,
	}
	//cors
	server.engine.Use(cors.New())

	// Register routes
	server.RegisterRouter()

	return &server
}

func (s *Server) Run() {
	s.engine.Listen(s.Port)
}
