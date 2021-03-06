package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/richardsang2008/MathTestsGenerator/controllers"
	"github.com/spf13/viper"
	//_ "./docs" // docs is generated by Swag CLI, you have to import it
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
	username:=viper.Get("database.username")
	password:=viper.Get("database.password")
	database:=viper.Get("database.database")
	dbhost:=viper.Get("database.host")
	hostport:=viper.Get("server.port")
	release:=viper.Get("Release")
	logdir:=viper.Get("log.logdir")
	logfile:=viper.Get("log.logfile")
	if release =="DEBUG"{
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	//open database
	dbconnection:= fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",username,password,dbhost,database)
	db,err:=gorm.Open("mysql",dbconnection)
	if release == "DEBUG" {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

	if err != nil {
		fmt.Errorf("database connection error %s",err)
	}
	defer db.Close()
	// Set the router as the default one provided by Gin
	a := controllers.Routes{}
	routes:=a.NewRoutes(db)
	var router *gin.Engine
	if release =="DEBUG" {
		router = routes.InitializeRoutes(logdir.(string), logfile.(string), true)
	} else {
		router = routes.InitializeRoutes(logdir.(string), logfile.(string), false)
	}

	fmt.Println(fmt.Sprintf("Server is starting at http://localhost:%s",hostport))
	router.Run(fmt.Sprintf(":%s",hostport))
}
