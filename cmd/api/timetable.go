package main

import (
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

	timetableMap := make(map[string][]*data.Timetable)

	for _, item := range timetable {
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

	timetableMap := make(map[string][]*data.Timetable)

	for _, item := range timetable {
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

	timetableMap := make(map[string][]*data.Timetable)

	for _, item := range timetable {
		timetableMap[item.ClasstimeDay] = append(timetableMap[item.ClasstimeDay], item)
	}

	c.JSON(http.StatusOK, gin.H{"payload":timetableMap})
	return
}
