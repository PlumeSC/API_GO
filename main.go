package main

import (
	"false_api/modules"
	adapterapi "false_api/modules/adapters/adapter_api"
	adapterauth "false_api/modules/adapters/adapter_auth"
	coreapi "false_api/modules/core/core_api"
	coreauth "false_api/modules/core/core_auth"
	"os"
)

func main() {
	app, db := modules.Init()

	authAdapter := adapterauth.NewAuthRepository(db)
	authService := coreauth.NewAuthService(authAdapter)
	authHandler := adapterauth.NewAuthHandler(authService)
	apiAdapter := adapterapi.NewApiRepository(db)
	apiRequest := adapterapi.NewApiRequest()
	apiService := coreapi.NewApiService(apiAdapter, apiRequest)
	apiHandler := adapterapi.NewApiHandler(apiService)

	app.Post("/register", authHandler.Register)
	app.Post("/Login", authHandler.Login)
	// app.Post("/createtable", apiHandler.CreateTables)    // league,season
	app.Post("/createplayers", apiHandler.CreaatePlayer) // league,season
	app.Post("/creatematch", apiHandler.CreateMatch)     // league,season,round

	app.Listen(os.Getenv("URL"))
}
