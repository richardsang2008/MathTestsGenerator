package controllers

import (
	"github.com/akath19/gin-zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"time"
)

type Routes struct {
}

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"logs/math.log",
	}
	return cfg.Build()
}
func (r *Routes) InitializeRoutes() *gin.Engine {
	router := gin.New()
	//Logging to a file
	zlog, _ := NewLogger()
	defer zlog.Sync()
	//Add middleware to Gin, requires sync duration & zap pointer
	router.Use(ginzap.Logger(3*time.Second, zlog))

	router.Use(gin.Recovery())
	studentController := StudentController{}
	quizController := QuizController{}
	pinController := PinController{}

	zlog.Info("Start the router")
	router.GET("/ping", pinController.Pinhandler)
	api := router.Group("/api")
	{
		studentapi := api.Group("/Student")
		{
			studentapi.POST("", studentController.CreateStudent)
			studentapi.GET("/byStudentId", studentController.GetStudentByStudentId)
			studentapi.GET("/byEmail", studentController.GetStudentByEmail)
		}
		quizapi := api.Group("/api/Quiz")
		{
			quizapi.GET("/:id", quizController.GetQuizById)
			quizapi.GET("/:id/score", quizController.GetQuizScoreById)
			quizapi.PATCH("/quizitems", quizController.PatchQuizItems)
		}
	}
	return router
}
