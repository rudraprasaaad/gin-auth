package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rudraprasaaad/gin-auth/controllers"
	"github.com/rudraprasaaad/gin-auth/middleware"
)

func UserRoutes(routes *gin.Engine) {
	routes.Use(middleware.Authenticate())
	routes.GET("/users", controller.GetUsers())
	routes.GET("/users/:user_id", controller.GetUser())
}
