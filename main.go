package main

import (
	"worker/database"
	"worker/routes"
	"worker/seeders"
)

func main() {
	database.Init()
	defer database.CloseDB()

	seeders.Seed()

	r := routes.SetupRouter()

	r.Static("/public", "./public")
	r.Run(":8080")
}
