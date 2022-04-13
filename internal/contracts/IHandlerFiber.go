package contracts

import "github.com/gofiber/fiber/v2"

type IHandlerFiber interface {
	Index(ctx *fiber.Ctx) error
	TopSecret(ctx *fiber.Ctx) error
	TopScretSplitPost(ctx *fiber.Ctx) error
	TopScretSplitGet(ctx *fiber.Ctx) error
}
