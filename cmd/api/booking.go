package main

import (
	"gin-api-template/internal/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
	id := c.Param("reserverId")

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

func (app *application) GetDateTimeBooking(c *gin.Context) {
	var input data.Booking

	if err := c.BindJSON(&input); err != nil {
		app.serverErrorResponse(err, c)
	}

	var noRooms []string

	inputStart, _ := time.Parse("15:04", input.StartTime)
	inputEnd, _ := time.Parse("15:04", input.EndTime)

	timetable, err := app.models.Timetables.GetByWeekDay("d"+strconv.Itoa(int(input.Date.Weekday())))
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			//app.NotFoundResponse(err, c)
			//return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	for _, item := range timetable {
		start, _ := time.Parse("15:04", item.ClasstimeTime)
		end := start.Add(time.Minute*50)

		if ((inputStart.After(start) || inputStart.Equal(start)) && inputStart.Before(end)) || (inputEnd.After(start) && (inputEnd.Before(end) || inputEnd.Equal(end))) || (inputStart.Before(start) && inputEnd.After(end)) {
			noRooms = append(noRooms, strconv.Itoa(int(item.RoomId)))
		}
	}

	booking, err := app.models.Booking.GetAllByDate(input.Date)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			//app.NotFoundResponse(err, c)
			//return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	for _, item := range booking {
		start, _ := time.Parse("15:04", item.StartTime)
		end, _ := time.Parse("15:04", item.EndTime)

		if ((inputStart.After(start) || inputStart.Equal(start)) && inputStart.Before(end)) || (inputEnd.After(start) && (inputEnd.Before(end) || inputEnd.Equal(end))) || (inputStart.Before(start) && inputEnd.After(end)) {
			noRooms = append(noRooms, strconv.Itoa(int(item.RoomId)))
		}
	}

	rooms, err := app.models.Extras.GetFreeRooms(strings.Join(noRooms, ","))
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"payload":rooms})
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

	inputStart, _ := time.Parse("15:04", input.StartTime)
	inputEnd, _ := time.Parse("15:04", input.EndTime)

	timetable, err := app.models.Timetables.GetByWeekDay("d"+strconv.Itoa(int(input.Date.Weekday())))
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			//app.NotFoundResponse(err, c)
			//return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	for _, item := range timetable {
		start, _ := time.Parse("15:04", item.ClasstimeTime)
		end := start.Add(time.Minute*50)

		if ((inputStart.After(start) || inputStart.Equal(start)) && inputStart.Before(end)) || (inputEnd.After(start) && (inputEnd.Before(end) || inputEnd.Equal(end))) || (inputStart.Before(start) && inputEnd.After(end)) {
			app.BadRequest(nil, c)
			return
		}
	}

	roomBooking, err := app.models.Booking.GetAllByRoom(input.RoomId)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{

		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	for _, item := range roomBooking {
		start, _ := time.Parse("15:04", item.StartTime)
		end, _ := time.Parse("15:04", item.EndTime)

		if ((inputStart.After(start) || inputStart.Equal(start)) && inputStart.Before(end)) || (inputEnd.After(start) && (inputEnd.Before(end) || inputEnd.Equal(end))) || (inputStart.Before(start) && inputEnd.After(end)) {
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
