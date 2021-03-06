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

	tt := router.Group("/tt")
	{
		tt.GET("/group", app.GetAllGroups)
		tt.GET("/teacher", app.GetAllTeachers)
		tt.GET("/room", app.GetAllRooms)
		tt.GET("/room/:roomId", app.GetRoom)

		timetable := tt.Group("/timetable")
		{
			timetable.GET("/group/:group", app.GetGroupTimetable)
			timetable.GET("/tutor/:tutorId", app.GetTutorTimetable)
			timetable.GET("/tutor/email/:email", app.GetTutorEmailTimetable)
			timetable.GET("/room/:roomId", app.GetRoomTimetable)
		}

		booking := tt.Group("/booking")
		{
			booking.GET("", app.GetAllBooking)
			booking.POST("/datetime", app.GetDateTimeBooking)

			booking.POST("/room/:roomId", app.GetRoomBooking)
			booking.GET("/reserver/:reserverId", app.GetReserverBooking)
			booking.POST("/reserver/between/:reserverId", app.GetReserverBetweenBooking)

			booking.GET("/requests", app.GetBookingRequests)
			booking.GET("/confirm", app.GetConfirmedBooking)

			secured := booking.Group("", app.auth())
			{
				secured.POST("/create", app.BookRoom)
				secured.POST("/confirm/:bookingId", app.ConfirmBooking)
				secured.POST("/reject/:bookingId", app.RejectBooking)
			}
		}

		login := tt.Group("/login")
		{
			login.POST("/office", app.OfficeAuth)
		}
		tt.GET("/isAdmin/:id", app.auth(), app.isAdmin)
	}

	//users.GET("/userdata", app.auth(), app.UserData)

	files := router.Group("/files")
	{
		files.POST("/upload/:fileName", app.FileUpload)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
