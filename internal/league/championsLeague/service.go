package championsLeague

import "errors"

type ChampionsLeagueService struct {
	leagueName string
}

func NewChampionsLeagueService() *ChampionsLeagueService {
	return &ChampionsLeagueService{leagueName: "cl"}
}
func (s ChampionsLeagueService) GetChampions(fairly int) (string, error) {
	if fairly == 0 {
		return "", errors.New("error")
	}
	if fairly > 5 {
		return "real madrid", nil
	} else if fairly == 5 {
		return "barcelona", nil
	} else {
		return "kiev", nil
	}
}
