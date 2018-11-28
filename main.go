package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richardsang2008/MathTestsGenerator/controllers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//db,err:=gorm.Open()

	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	// Set the router as the default one provided by Gin
	routes := controllers.Routes{}
	router := routes.InitializeRoutes()
	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
