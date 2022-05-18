package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "gin-api-template/cmd/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"
)

func (app *application) routes() http.Handler {

	//this use logger and recovery middleware by default, use in dev mode.
	router := gin.Default()

	//this has no logger and recovery, so include it in middleware list if needed.
	//router := gin.New()


	//if we need to serve static uploads or return html use following
	/*
	router.LoadHTMLGlob("./ui/*.html")
	router.Static("/api/serve", "./ui/")
	*/


	//list middleware that u want to include by default
	router.Use(
		//enabling AllowAllOrigins = true
		cors.Default(),

		//include in prod mode
		//gin.Recovery(),
		)

	router.GET("/healthcheck", app.Healthcheck)

	router.GET("/getTimetable/group/:groupId", app.GetGroupTimetable)
	router.GET("/getTimetable/tutor/:tutorId", app.GetTutorTimetable)
	router.GET("/getTimetable/room/:roomId", app.GetRoomTimetable)

	users := router.Group("/users")
	{
		login := users.Group("/login")
		{
			login.POST("/email", app.EmailAuth)
		}

		register := users.Group("/register")
		{
			register.POST("/new", app.UserCreate)
		}

		users.GET("/userdata", app.auth(), app.UserData)
	}

	files := router.Group("/files")
	{
		files.POST("/upload/:fileName", app.FileUpload)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}