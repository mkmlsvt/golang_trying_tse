package concretes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sony/gobreaker"
	"golangTdd/internal/interfaces"
	"strconv"
	"time"
)

type SampiyonlarLigiEndpoint struct {
	service        interfaces.LeagueService
	name           string
	circuitBreaker *gobreaker.CircuitBreaker
}

func NewSampiyonlarLigiEndpoint(service interfaces.LeagueService, name string) *SampiyonlarLigiEndpoint {
	settings := gobreaker.Settings{
		Name:        "SampiyonlarLigiService",
		MaxRequests: 3,
		Interval:    60 * time.Second,
		Timeout:     10 * time.Second,
	}
	cb := gobreaker.NewCircuitBreaker(settings)
	return &SampiyonlarLigiEndpoint{service: service, name: name, circuitBreaker: cb}
}

func (s SampiyonlarLigiEndpoint) Name() string {
	return s.name
}

// cl godoc
// @Summary get cl league
// @Description get cl league
// @Tags League
// @Accept json
// @Produce json
// @Param fair query int false "fair"
// @Success 200
// @Router /leagues/champion/cl/ [get]
func (s SampiyonlarLigiEndpoint) Get() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fairly, _ := strconv.Atoi(ctx.Query("fair"))
		wrappedFunction := func() (interface{}, error) {
			return s.service.GetChampions(fairly)
		}
		result, err := s.circuitBreaker.Execute(wrappedFunction)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return ctx.JSON(result)
	}
}
