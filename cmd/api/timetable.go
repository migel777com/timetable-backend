package main

import (
	"gin-api-template/internal/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (app *application) GetTimetable(c *gin.Context){
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	timetable, err := app.models.Timetables.Get(id)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.InvalidCredentials(err, c)
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
