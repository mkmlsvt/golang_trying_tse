package integrationTests

import (
	"golangTdd/api"
	"golangTdd/api/endpoints/concretes"
	"golangTdd/api/endpoints/interfaces"
	"golangTdd/internal/league/championsLeague"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChampionTeamEndpoint(t *testing.T) {
	clService := championsLeague.NewChampionsLeagueService()
	endpoint := concretes.NewSampiyonlarLigiEndpoint(clService, "cl")
	app := api.NewServer([]interfaces.Endpoint{endpoint})
	testServer := app.SetupApp()
	t.Run("should return real madrid", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/leagues/champion/cl?fair=10", nil)
		resp, err := testServer.Test(req)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		actualChampionTeam := string(bodyBytes)
		expectedChampionTeam := "\"real madrid\""

		assert.Equal(t, expectedChampionTeam, actualChampionTeam, "Beklenen şampiyon takımı")
	})

}
