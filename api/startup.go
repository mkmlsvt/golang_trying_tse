package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "golangTdd/api/docs"
	"golangTdd/api/endpoints/interfaces"
)

//go:generate swag init -g startup.go

type server struct {
	endPoints []interfaces.Endpoint
}

type Server interface {
	SetupApp() *fiber.App
}

func NewServer(
	endPoints []interfaces.Endpoint,

) Server {
	return &server{
		endPoints: endPoints,
	}
}

func (s server) SetupApp() *fiber.App {
	app := fiber.New()
	group := app.Group("/leagues/champion")
	for _, endPoint := range s.endPoints {
		group.Get("/"+endPoint.Name(), endPoint.Get())
	}
	app.Get("/*", swagger.HandlerDefault) // default

	return app
}
