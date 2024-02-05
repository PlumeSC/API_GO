package main

import (
	"false_api/modules"
	"os"
)

func main() {
	app, db := modules.Init()
	_ = db

	app.Listen(os.Getenv("URL"))
}
