package main

import (
	"context"
	"golangTdd/api"
	"golangTdd/api/endpoints/concretes"
	"golangTdd/api/endpoints/interfaces"
	"golangTdd/internal/league/championsLeague"
	"golangTdd/internal/league/superLeauge"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	superligService := superLeauge.NewSuperLeagueService()
	clService := championsLeague.NewChampionsLeagueService()
	clEnpoint := concretes.NewSampiyonlarLigiEndpoint(clService, "cl")
	stslEnpoint := concretes.NewSuperLigEndpoint(superligService, "stsl")

	router := api.NewServer([]interfaces.Endpoint{clEnpoint, stslEnpoint})
	httpClient := api.NewHttpServer(":8080", router.SetupApp())
	httpClient.Start(context.Background(), wg)
	wg.Wait()
}
