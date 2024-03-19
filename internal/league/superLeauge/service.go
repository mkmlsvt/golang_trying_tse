package superLeauge

type SuperLeagueService struct {
	leagueName string
}

func NewSuperLeagueService() *SuperLeagueService {
	return &SuperLeagueService{leagueName: "stsl"}
}
func (s SuperLeagueService) GetChampions(fairly int) string {
	if fairly > 5 {
		return "fenerbahçe"
	} else {
		return "otherTeams"
	}
}
