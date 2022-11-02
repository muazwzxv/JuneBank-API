package main

import (
	_ "github.com/jackc/pgx/v5"
)

func main() {

	var app Application

	app.setup()
	app.registerRoutes()
	app.run()
}
