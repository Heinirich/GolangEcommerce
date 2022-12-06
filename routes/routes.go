package routes

import (
	"github.com/heinrich/golangecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func userRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.POST("users/signup")
	incomingRoutes.POST("users/login")
	incomingRoutes.POST("/admin/addproduct")
	incomingRoutes.GET("/users/productreview")
	incomingRoutes.GET("/users/search")
}