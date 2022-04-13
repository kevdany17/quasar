package handlers

import (
	"quasar/internal/contracts"

	"github.com/gofiber/fiber/v2"
)

type HandlerImpl struct {
	service contracts.IServiceSatellite
}

func NewHandlerImpl(s contracts.IServiceSatellite) contracts.IHandlerFiber {
	return &HandlerImpl{
		service: s,
	}
}

func (h HandlerImpl) Index(c *fiber.Ctx) error {
	return c.SendString("Up Server!")
}

func (h HandlerImpl) TopSecret(c *fiber.Ctx) error {
	return h.service.ResponseForList(c)
}

func (h HandlerImpl) TopScretSplitPost(c *fiber.Ctx) error {
	return h.service.RequestForName(c)
}

func (h HandlerImpl) TopScretSplitGet(c *fiber.Ctx) error {
	return h.service.ResponseForName(c)
}
