package main

import (
	"false_api/modules"

	authadapter "false_api/modules/adapters/auth_adapters"
	matchsadapters "false_api/modules/adapters/matchs_adapters"
	seasonadapters "false_api/modules/adapters/season_adapters"
	standingsadapters "false_api/modules/adapters/standings_adapters"
	authcore "false_api/modules/core/auth_core"
	matchscore "false_api/modules/core/matchs_core"
	seasoncore "false_api/modules/core/season_core"
	standingscore "false_api/modules/core/standings_core"

	"os"
)

func main() {
	app, db := modules.Init()

	authAdapter := authadapter.NewAuthRepository(db)
	authService := authcore.NewAuthService(authAdapter)
	authHandler := authadapter.NewAuthHandler(authService)
	seasonAdapter := seasonadapters.NewSeasonRepository(db)
	seasonApi := seasonadapters.NewApiFootball()
	seasonService := seasoncore.NewSeasonService(seasonAdapter, seasonApi)
	seasonHandler := seasonadapters.NewSeasonHandler(seasonService)
	standingsRepository := standingsadapters.NewstandingsRepository(db)
	standingsService := standingscore.NewStandingsService(standingsRepository, seasonApi)
	standingsHandler := standingsadapters.NewStandingsHandler(standingsService)
	matchsRepository := matchsadapters.NewMatchsRepository(db)
	matchsService := matchscore.NewMatchService(matchsRepository, seasonApi)
	matchsHandler := matchsadapters.NewMatchsHandler(matchsService)

	app.Post("/register", authHandler.Register)
	app.Post("/Login", authHandler.Login)
	app.Post("/createseason", seasonHandler.CreateStandings)      // league season
	app.Post("/createplayer", seasonHandler.CreatePlayers)        // league season
	app.Post("/creatematch", seasonHandler.CreateMatch)           // league season
	app.Get("/getstandings", standingsHandler.GetStandings)       // league season
	app.Get("/updatestandings", standingsHandler.UpdateStandings) // league season
	app.Get("/getmatches", matchsHandler.GetAll)                  // teamName round? league season
	app.Get("/updatematches", matchsHandler.UpdateMatch)          // round league season

	app.Listen(os.Getenv("URL"))
}
