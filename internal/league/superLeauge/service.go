package superLeauge

import "errors"

type SuperLeagueService struct {
	leagueName string
}

func NewSuperLeagueService() *SuperLeagueService {
	return &SuperLeagueService{leagueName: "stsl"}
}
func (s SuperLeagueService) GetChampions(fairly int) (string, error) {
	if fairly == 0 {
		return "", errors.New("errorrrr")
	}
	if fairly > 5 {
		return "fenerbahçe", nil
	} else {
		return "otherTeams", nil
	}
}
