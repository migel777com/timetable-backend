package main

import (
	"encoding/json"
	"gin-api-template/internal/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (app *application) GetGroupTimetable(c *gin.Context){
	group := c.Param("group")

	timetable, err := app.models.Timetables.GetByGroup(group)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		}
		app.serverErrorResponse(err, c)
		return
	}

	days, err := app.models.Timetables.GetDateTimeData(timetable[0].ScheduleBlockId)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotAcceptable(err, c)
			return
		}
		app.serverErrorResponse(err, c)
		return
	}

	dateTimeData := data.DateTimeData{}
	err = json.Unmarshal([]byte(days), &dateTimeData)
	if err != nil {
		app.serverErrorResponse(err, c)
		return
	}

	timeMap := make(map[string]*data.Time)

	for _, time := range dateTimeData.Time {
		timeMap[time.Id] = time
	}

	timetableMap := make(map[string][]*data.Timetable)

	for _, item := range timetable {
		item.StartTime = timeMap[item.ClasstimeTime].Start
		item.EndTime = timeMap[item.ClasstimeTime].Finish
		timetableMap[item.ClasstimeDay] = append(timetableMap[item.ClasstimeDay], item)
	}

	c.JSON(http.StatusOK, gin.H{"payload":timetableMap})
	return
}

func (app *application) GetTutorTimetable(c *gin.Context){
	id, _ := strconv.ParseInt(c.Param("tutorId"), 10, 64)

	timetable, err := app.models.Timetables.GetByTutor(id)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		}
		app.serverErrorResponse(err, c)
		return
	}

	days, err := app.models.Timetables.GetDateTimeData(timetable[0].ScheduleBlockId)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotAcceptable(err, c)
			return
		}
		app.serverErrorResponse(err, c)
		return
	}

	dateTimeData := data.DateTimeData{}
	err = json.Unmarshal([]byte(days), &dateTimeData)
	if err != nil {
		app.serverErrorResponse(err, c)
		return
	}

	timeMap := make(map[string]*data.Time)

	for _, time := range dateTimeData.Time {
		timeMap[time.Id] = time
	}

	timetableMap := make(map[string][]*data.Timetable)

	for _, item := range timetable {
		item.StartTime = timeMap[item.ClasstimeTime].Start
		item.EndTime = timeMap[item.ClasstimeTime].Finish
		timetableMap[item.ClasstimeDay] = append(timetableMap[item.ClasstimeDay], item)
	}

	c.JSON(http.StatusOK, gin.H{"payload":timetableMap})
	return
}

func (app *application) GetTutorEmailTimetable(c *gin.Context){
	email := c.Param("email")

	timetable, err := app.models.Timetables.GetByTutorEmail(email)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		}
		app.serverErrorResponse(err, c)
		return
	}

	days, err := app.models.Timetables.GetDateTimeData(timetable[0].ScheduleBlockId)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotAcceptable(err, c)
			return
		}
		app.serverErrorResponse(err, c)
		return
	}

	dateTimeData := data.DateTimeData{}
	err = json.Unmarshal([]byte(days), &dateTimeData)
	if err != nil {
		app.serverErrorResponse(err, c)
		return
	}

	timeMap := make(map[string]*data.Time)

	for _, time := range dateTimeData.Time {
		timeMap[time.Id] = time
	}

	timetableMap := make(map[string][]*data.Timetable)

	for _, item := range timetable {
		item.StartTime = timeMap[item.ClasstimeTime].Start
		item.EndTime = timeMap[item.ClasstimeTime].Finish
		timetableMap[item.ClasstimeDay] = append(timetableMap[item.ClasstimeDay], item)
	}

	c.JSON(http.StatusOK, gin.H{"payload":timetableMap})
	return
}

func (app *application) GetRoomTimetable(c *gin.Context){
	id, _ := strconv.ParseInt(c.Param("roomId"), 10, 64)

	timetable, err := app.models.Timetables.GetByRoom(id)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		}
		app.serverErrorResponse(err, c)
		return
	}

	days, err := app.models.Timetables.GetDateTimeData(timetable[0].ScheduleBlockId)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotAcceptable(err, c)
			return
		}
		app.serverErrorResponse(err, c)
		return
	}

	dateTimeData := data.DateTimeData{}
	err = json.Unmarshal([]byte(days), &dateTimeData)
	if err != nil {
		app.serverErrorResponse(err, c)
		return
	}

	timeMap := make(map[string]*data.Time)

	for _, time := range dateTimeData.Time {
		timeMap[time.Id] = time
	}

	timetableMap := make(map[string][]*data.Timetable)

	for _, item := range timetable {
		item.StartTime = timeMap[item.ClasstimeTime].Start
		item.EndTime = timeMap[item.ClasstimeTime].Finish
		timetableMap[item.ClasstimeDay] = append(timetableMap[item.ClasstimeDay], item)
	}

	c.JSON(http.StatusOK, gin.H{"payload":timetableMap})
	return
}
