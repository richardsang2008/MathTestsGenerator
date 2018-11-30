package controllers

import (
	"fmt"
	"github.com/akath19/gin-zap"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/wantedly/gorm-zap"
	"go.uber.org/zap"
	"os"

	"time"
)

type Routes struct {
	Db *gorm.DB
}
func (r *Routes) NewRoutes(l *gorm.DB) *Routes {
	r.Db =l
	return r
}
func newLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	logdir :="logs"
	logfile:="math.log"
	var _, err = os.Stat(logdir)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(logdir,0755)
	}
	logpath:=fmt.Sprintf("%s/%s",logdir,logfile)
	cfg.OutputPaths = []string{
		logpath,
	}
	return cfg.Build()
}


func (r *Routes) InitializeRoutes() *gin.Engine {
	router := gin.New()
	//use ginSwagger middleware
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Static("/swaggerui/", "cmd/api/swaggerui")
	//Logging to a file
	zlog, _ := newLogger()
	defer zlog.Sync()
	//Add middleware to Gin, requires sync duration & zap pointer
	router.Use(ginzap.Logger(3*time.Second, zlog))

	router.Use(gin.Recovery())
	a := StudentController{}
	studentController:=a.NewStudentController(r.Db)
	b := QuizController{}
	quizController := b.NewQuizController(r.Db)

	pinController := PinController{}

	zlog.Info("Start the router")
	r.Db.SetLogger(gormzap.New(zlog))
	router.GET("/api/ping", pinController.Pinhandler)
	router.POST("/api/Student",studentController.CreateStudent)
	api := router.Group("/api")
	{
		studentapi := api.Group("/Student")
		{
			//studentapi.POST("", studentController.CreateStudent)
			studentapi.GET("/byStudentId", studentController.GetStudentByStudentId)
			studentapi.GET("/byEmail", studentController.GetStudentByEmail)
		}
		api.POST("/Quiz",quizController.CreateAQuiz)
		quizapi := api.Group("/Quiz")
		{
			quizapi.GET("/:id", quizController.GetAQuizById)
			quizapi.GET("/:id/score", quizController.ScoreTheQuiz)
			quizapi.PATCH("/quizitems", quizController.AnswerAQuizItem)

		}
	}
	return router
}
