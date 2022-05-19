package main

import (
	"gin-api-template/internal/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
	booking, err := app.models.Booking.GetAllRequests()
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

	inputStart, _ := time.Parse("15:04", input.StartTime)
	inputEnd, _ := time.Parse("15:04", input.EndTime)

	for _, item := range roomBooking {
		start, _ := time.Parse("15:04", item.StartTime)
		end, _ := time.Parse("15:04", item.EndTime)

		if (inputStart.After(start) && inputStart.Before(end)) || (inputEnd.After(start) && inputEnd.Before(end)) || (inputStart.Before(start) && inputEnd.After(end)) {
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
