package concretes

import (
	"github.com/gofiber/fiber/v2"
	"golangTdd/internal/interfaces"
	"strconv"
)

type SuperLigEndpoint struct {
	service interfaces.LeagueService
	name    string
}

func NewSuperLigEndpoint(service interfaces.LeagueService, name string) *SuperLigEndpoint {
	return &SuperLigEndpoint{service: service, name: name}
}

func (s SuperLigEndpoint) Name() string {
	return s.name
}

// stsl godoc
// @Summary get stsl league
// @Description get stsl league
// @Tags League
// @Accept json
// @Produce json
// @Param fair query int false "fair"
// @Success 200 {string} string "Grouped widget data"
// @Router /leagues/champion/stsl/ [get]
func (s SuperLigEndpoint) Get() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fairly, _ := strconv.Atoi(ctx.Params("fair"))
		championTeam := s.service.GetChampions(fairly)
		return ctx.JSON(championTeam)
	}
}
