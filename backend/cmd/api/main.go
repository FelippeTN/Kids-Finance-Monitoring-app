package main

import (
	"github.com/FelippeTN/Kids-Finance-Monitoring-app/backend/config"
	"github.com/FelippeTN/Kids-Finance-Monitoring-app/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

//TODO: Criar um swagger UI para documentar a API com swaggo/swag
func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.Run(":8000")
}