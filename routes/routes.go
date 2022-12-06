package routes

import (
	"github.com/heinrich/golangecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.POST("users/signup",controllers.Signup)
	incomingRoutes.POST("users/login",controllers.Login)
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productreview",controllers.SearchhProduct())
	incomingRoutes.GET("/users/search",controllers.SearchProductByQuery)
}