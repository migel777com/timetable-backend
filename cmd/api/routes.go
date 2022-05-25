package main

import (
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
		//cors.Default(),
		CORSMiddleware(),

		//include in prod mode
		//gin.Recovery(),
		)

	router.GET("/healthcheck", app.Healthcheck)

	router.GET("/group", app.GetAllGroups)
	router.GET("/teacher", app.GetAllTeachers)
	router.GET("/room", app.GetAllRooms)
	router.GET("/room/:roomId", app.GetRoom)

	timetable := router.Group("timetable")
	{
		timetable.GET("/group/:group", app.GetGroupTimetable)
		timetable.GET("/tutor/:tutorId", app.GetTutorTimetable)
		timetable.GET("/room/:roomId", app.GetRoomTimetable)
	}

	booking := router.Group("/booking")
	{
		booking.GET("", app.GetAllBooking)
		booking.POST("/datetime", app.GetDateTimeBooking)

		booking.POST("/room/:roomId", app.GetRoomBooking)
		booking.GET("/reserver/:reserverId", app.GetReserverBooking)

		booking.GET("/requests", app.GetBookingRequests)

		secured := booking.Group("", app.auth())
		{
			secured.POST("/create", app.BookRoom)
			secured.PATCH("/confirm/:bookingId", app.ConfirmBooking)
			secured.PATCH("/reject/:bookingId", app.RejectBooking)
		}
	}

	login := router.Group("/login")
	{
		login.POST("/office", app.OfficeAuth)
	}
	//users.GET("/userdata", app.auth(), app.UserData)


	files := router.Group("/files")
	{
		files.POST("/upload/:fileName", app.FileUpload)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
