package services

import (
	"quasar/internal/contracts"
	"quasar/internal/models"

	"github.com/gofiber/fiber/v2"
)

type ServiceSatelliteImpl struct {
	quasarService contracts.IQuasar
	repository    contracts.IRepository
}

func NewServiceSatelliteImpl(repository contracts.IRepository, quasar contracts.IQuasar) contracts.IServiceSatellite {
	return &ServiceSatelliteImpl{
		repository:    repository,
		quasarService: quasar,
	}
}
func (s *ServiceSatelliteImpl) ResponseForList(ctx *fiber.Ctx) error {
	var request models.Request

	err := ctx.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
	var distances []float32
	var messages [][]string
	for _, satellite := range request.Satellites {
		distances = append(distances, satellite.Distance)
	}
	for _, satellite := range request.Satellites {
		messages = append(messages, satellite.Message)
	}
	x, y := s.quasarService.GetLocation(distances...)

	message := s.quasarService.GetMessage(messages...)

	if (x == 0 && y == 0) || message == "" {
		return fiber.NewError(404, "Position not found")
	}

	response := models.SatelliteResponse{
		Position: models.Position{
			X: x,
			Y: y,
		},
		Message: message,
	}

	return ctx.Status(200).JSON(response)
}
func (s *ServiceSatelliteImpl) RequestForName(ctx *fiber.Ctx) error {
	var body models.SatelliteRequest

	err := ctx.BodyParser(&body)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	body.Name = ctx.Params("satilliteName", "")

	if body.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Name Satellite is required")
	}

	x, y := s.quasarService.GetLocation(body.Distance)

	message := s.quasarService.GetMessage(body.Message)

	if (x == 0 && y == 0) || message == "" {
		return fiber.NewError(404, "Position not found")
	}

	storage := models.Storage{
		Satellite: body,
		Position: models.Position{
			X: x,
			Y: y,
		},
		Message: message,
	}

	err = s.repository.Save(storage)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	return ctx.Status(201).SendString("Satellite created")
}
func (s *ServiceSatelliteImpl) ResponseForName(ctx *fiber.Ctx) error {
	var name string

	name = ctx.Params("satilliteName", "")

	if name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Name Satellite is required")
	}

	storage, err := s.repository.GetForName(name)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Satillite not found")
	}

	response := models.SatelliteResponse{
		Position: storage.Position,
		Message:  storage.Message,
	}

	return ctx.Status(200).JSON(response)
}
