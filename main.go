package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richardsang2008/MathTestsGenerator/controllers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"fmt"
)

func main() {
	//db,err:=gorm.Open()
	//appconfig loading
	viper.SetConfigName("appconfig")
	viper.AddConfigPath("config")
	viper.ReadInConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:",e.Name)
	})

	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	// Set the router as the default one provided by Gin
	routes := controllers.Routes{}
	router := routes.InitializeRoutes()
	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
