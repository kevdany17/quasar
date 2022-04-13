package contracts

import "github.com/gofiber/fiber/v2"

type IServiceSatellite interface {
	ResponseForList(ctx *fiber.Ctx) error
	RequestForName(ctx *fiber.Ctx) error
	ResponseForName(ctx *fiber.Ctx) error
}
