package main

import (
	"false_api/modules"
	adaptersauth "false_api/modules/adapters/adapters_auth"
	coreauth "false_api/modules/core/core_auth"
	"os"
)

func main() {
	app, db := modules.Init()

	authAdapter := adaptersauth.NewAuthRepository(db)
	authService := coreauth.NewAuthService(authAdapter)
	authHandler := adaptersauth.NewAuthHandler(authService)

	app.Post("/register", authHandler.Register)

	app.Listen(os.Getenv("URL"))
}
