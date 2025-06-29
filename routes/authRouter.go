package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rudraprasaaad/gin-auth/controllers"
)

func AuthRoutes(routes *gin.Engine) {
	routes.POST("users/signup", controller.Signup())
	routes.POST("users/login", controller.Login())
}
