package championsLeague

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ClChampions(t *testing.T) {
	clService := NewChampionsLeagueService()
	var result string
	var err error
	t.Run("Check if fair gt 5", func(t *testing.T) {
		result, err = clService.GetChampions(6)
		assert.Equal(t, "real madrid", result)
		assert.Nil(t, err)
		result, err = clService.GetChampions(10)
		assert.Equal(t, "real madrid", result)
		assert.Nil(t, err)
	})
	t.Run("Check if fair lt 5", func(t *testing.T) {
		result, err = clService.GetChampions(1)
		assert.Equal(t, "kiev", result)
		assert.Nil(t, err)
		result, err = clService.GetChampions(3)
		assert.Equal(t, "kiev", result)
		assert.Nil(t, err)
	})
	t.Run("Check if fair equal 5", func(t *testing.T) {
		result, err = clService.GetChampions(5)
		assert.Equal(t, "barcelona", result)
		assert.Nil(t, err)
	})
	t.Run("Check if fair equal 0", func(t *testing.T) {
		result, err = clService.GetChampions(0)
		assert.Equal(t, "", result)
		assert.NotNil(t, err)
	})
}
