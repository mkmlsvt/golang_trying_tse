package interfaces

import "github.com/gofiber/fiber/v2"

type Endpoint interface {
	Name() string
	Get() fiber.Handler
}
