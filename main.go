package main

import (
	"false_api/modules"

	authadapter "false_api/modules/adapters/auth_adapters"
	seasonadapters "false_api/modules/adapters/season_adapters"
	authcore "false_api/modules/core/auth_core"
	seasoncore "false_api/modules/core/season_core"

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

	app.Post("/register", authHandler.Register)
	app.Post("/Login", authHandler.Login)
	app.Post("/createseason", seasonHandler.CreateStandings) // league season
	app.Post("/createplayer", seasonHandler.CreatePlayers)   // league season
	app.Post("/creatematch", seasonHandler.CreateMatch)      // league season


	app.Listen(os.Getenv("URL"))
}
