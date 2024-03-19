package concretes

import (
	"github.com/gofiber/fiber/v2"
	"golangTdd/internal/interfaces"
	"strconv"
)

type SampiyonlarLigiEndpoint struct {
	service interfaces.LeagueService
	name    string
}

func NewSampiyonlarLigiEndpoint(service interfaces.LeagueService, name string) *SampiyonlarLigiEndpoint {
	return &SampiyonlarLigiEndpoint{service: service, name: name}
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
		championTeam := s.service.GetChampions(fairly)
		return ctx.JSON(championTeam)
	}
}
