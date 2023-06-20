package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kala-111/gin-gorm-rest/config"
	"github.com/kala-111/gin-gorm-rest/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)

	router.Run(":8080")
}
