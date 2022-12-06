package  main

import (
	"github.com/heinrich/golangecommerce/controllers"
	"github.com/heinrich/golangecommerce/database"
	"github.com/heinrich/golangecommerce/middleware"
	"github.com/heinrich/golangecommerce/routes"

)

func main()  {
	port := os.Getenv("PORT")
	if port == ""{
		port = :8000
	}

	app := controllers.NewApplication(database.ProductData(database.Client,"Products"),database.UserData(database.Client,"Users"))

	router := gin.New()

}