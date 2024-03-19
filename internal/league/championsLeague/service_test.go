package championsLeague

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ClChampions(t *testing.T) {
	clService := NewChampionsLeagueService()
	t.Run("Check if fair gt 5", func(t *testing.T) {
		assert.Equal(t, "real madrid", clService.GetChampions(6))
		assert.Equal(t, "real madrid", clService.GetChampions(6))
		assert.Equal(t, "real madrid", clService.GetChampions(10))
	})
	t.Run("Check if fair lt 5", func(t *testing.T) {
		assert.Equal(t, "kiev", clService.GetChampions(1))
		assert.Equal(t, "kiev", clService.GetChampions(3))
	})
	t.Run("Check if fair equal 5", func(t *testing.T) {
		assert.Equal(t, "barcelona", clService.GetChampions(5))
	})
}
