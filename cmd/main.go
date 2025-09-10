package main

import (
	"auth/database"
	"auth/internal/configuration"
	"auth/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(configuration.Envs.GinMode)
	db := database.InitDB()
	r := routes.GetRoutes(db)

	r.Run(":" + configuration.Envs.Port)
}
