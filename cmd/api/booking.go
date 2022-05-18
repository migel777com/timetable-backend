package main

import (
	"gin-api-template/internal/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) GetAllBooking(c *gin.Context) {
	booking, err := app.models.Booking.GetAll()
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	bookingMap := make(map[string][]*data.Booking)

	for _, item := range booking {
		bookingMap[item.Day] = append(bookingMap[item.Day], item)
	}

	c.JSON(http.StatusOK, gin.H{"payload":bookingMap})
	return
}

func (app *application) GetRoomBooking(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("roomId"), 10, 64)

	booking, err := app.models.Booking.GetAllByRoom(id)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	bookingMap := make(map[string][]*data.Booking)

	for _, item := range booking {
		bookingMap[item.Day] = append(bookingMap[item.Day], item)
	}

	c.JSON(http.StatusOK, gin.H{"payload":bookingMap})
	return
}

func (app *application) GetReserverBooking(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("reserverId"), 10, 64)

	booking, err := app.models.Booking.GetAllByReserver(id)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	bookingMap := make(map[string][]*data.Booking)

	for _, item := range booking {
		bookingMap[item.Day] = append(bookingMap[item.Day], item)
	}

	c.JSON(http.StatusOK, gin.H{"payload":bookingMap})
	return
}

func (app *application) GetBookingRequests(c *gin.Context) {
	booking, err := app.models.Booking.GetAllConfirmed()
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	bookingMap := make(map[string][]*data.Booking)

	for _, item := range booking {
		bookingMap[item.Day] = append(bookingMap[item.Day], item)
	}

	c.JSON(http.StatusOK, gin.H{"payload":bookingMap})
	return
}

func (app *application) GetConfirmedBooking(c *gin.Context) {
	booking, err := app.models.Booking.GetAllConfirmed()
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	bookingMap := make(map[string][]*data.Booking)

	for _, item := range booking {
		bookingMap[item.Day] = append(bookingMap[item.Day], item)
	}

	c.JSON(http.StatusOK, gin.H{"payload":bookingMap})
	return
}

func (app *application) ConfirmBooking(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("bookingId"), 10, 64)

	err := app.models.Booking.Confirm(id)
	if err!=nil{
		if err.Error()=="no affected rows"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"payload":"booking confirmed"})
	return
}

func (app *application) RejectBooking(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("bookingId"), 10, 64)

	err := app.models.Booking.Delete(id)
	if err!=nil{
		if err.Error()=="no affected rows"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"payload":"booking rejected"})
	return
}

func (app *application) BookRoom(c *gin.Context) {
	var input data.Booking

	if err := c.BindJSON(&input); err != nil {
		app.serverErrorResponse(err, c)
	}

	roomBooking, err := app.models.Booking.GetAllByRoom(input.RoomId)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{

		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	start := strings.Split(input.StartTime, ":")
	end := strings.Split(input.EndTime, ":")

	inputStartH, _ := strconv.ParseInt(start[0], 10, 64)
	inputStartM, _ := strconv.ParseInt(start[1], 10, 64)

	inputEndH, _ := strconv.ParseInt(end[0], 10, 64)
	inputEndM, _ := strconv.ParseInt(end[1], 10, 64)

	for _, item := range roomBooking {
		start = strings.Split(item.StartTime, ":")
		end = strings.Split(item.EndTime, ":")

		startH, _ := strconv.ParseInt(start[0], 10, 64)
		startM, _ := strconv.ParseInt(start[1], 10, 64)

		endH, _ := strconv.ParseInt(end[0], 10, 64)
		endM, _ := strconv.ParseInt(end[1], 10, 64)

		if ((inputStartH >= startH && inputStartM >= startM) && (inputStartH < endH && inputStartM < endM)) || ((inputEndH >= startH && inputEndM >= startM) && (inputEndH < endH && inputEndM < endM)) {
			app.BadRequest(nil, c)
			return
		}
	}

	insert, err := app.models.Booking.Insert(input.Room, input.Reserver, input.ReserverInfo, input.Day, input.StartTime, input.EndTime, input.Reason, input.RoomId, input.ReserverId, input.Date)
	if err != nil {
		app.serverErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"payload":insert})
	return
}
