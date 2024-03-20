package main

import (
	"false_api/modules"

	authadapter "false_api/modules/adapters/auth_adapters"
	competitiveadapters "false_api/modules/adapters/competitive_adapters"
	matchsadapters "false_api/modules/adapters/matchs_adapters"
	seasonadapters "false_api/modules/adapters/season_adapters"
	standingsadapters "false_api/modules/adapters/standings_adapters"
	authcore "false_api/modules/core/auth_core"
	competitivecore "false_api/modules/core/competitive_core"
	matchscore "false_api/modules/core/matchs_core"
	seasoncore "false_api/modules/core/season_core"
	standingscore "false_api/modules/core/standings_core"

	"os"
)

var key string = os.Getenv("SECRET_KEY")

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

	compRepository := competitiveadapters.NewCompRepository(db)
	compService := competitivecore.NewCompService(compRepository, seasonApi, matchsRepository, matchsService, standingsService)
	compHandler := competitiveadapters.NewCompHandler(compService)

	compHandler.InitLive()

	jwt := modules.Middleware(key)

	app.Post("/register", authHandler.Register)
	app.Post("/Login", authHandler.Login)
	app.Post("/createseason", jwt, modules.Admin, seasonHandler.CreateStandings) // league season
	app.Post("/createplayer", jwt, modules.Admin, seasonHandler.CreatePlayers)   // league season
	app.Post("/creatematch", jwt, modules.Admin, seasonHandler.CreateMatch)      // league season

	app.Get("/getstandings", standingsHandler.GetStandings)       // league season
	app.Get("/updatestandings", standingsHandler.UpdateStandings) // league season
	app.Get("/player", matchsHandler.GetPlayer)
	app.Get("/getmatches", matchsHandler.GetAll)         // teamName round? league season
	app.Get("/updatematches", matchsHandler.UpdateMatch) // round league season
	app.Get("/comp", compHandler.GetMatchDay)

	app.Listen(os.Getenv("PORT"))
}
