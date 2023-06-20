package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kala-111/gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
