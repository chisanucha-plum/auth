package main

import (
	"auth/database"
	"auth/internal/configuration"
	"auth/internal/routes"
)

func main() {
	db := database.InitDB()
	r := routes.GetRoutes(db)

	r.Run(":" + configuration.Envs.Port)
}
