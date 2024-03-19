package championsLeague

type ChampionsLeagueService struct {
	leagueName string
}

func NewChampionsLeagueService() *ChampionsLeagueService {
	return &ChampionsLeagueService{leagueName: "cl"}
}
func (s ChampionsLeagueService) GetChampions(fairly int) string {
	if fairly > 5 {
		return "real madrid"
	} else if fairly == 5 {
		return "barcelona"
	} else {
		return "kiev"
	}
}
