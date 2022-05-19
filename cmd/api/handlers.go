package main

import (
	data "gin-api-template/internal/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Healthcheck GetUser godoc
// @Summary Retrieves application status
// @Produce json
// @Success 200 {object} data.Healthcheck
// @Router /healthcheck [get]
func (app *application) Healthcheck(c *gin.Context) {

	health := &data.Healthcheck{
		Status:      "available",
		Environment: app.config.Env,
		Version:     version,
	}

	c.JSON(http.StatusOK, gin.H{"payload":health})

}


func (app *application) GetAllGroups(c *gin.Context) {
	groups, err := app.models.Extras.GetAllGroups()
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"payload":groups})
	return
}

func (app *application) GetAllTeachers(c *gin.Context) {
	teachers, err := app.models.Extras.GetAllTeachers()
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"payload":teachers})
	return
}

func (app *application) GetAllRooms(c *gin.Context) {
	rooms, err := app.models.Extras.GetAllRooms()
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

func (app *application) GetRoom(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("roomId"), 10, 64)

	room, err := app.models.Extras.GetRoom(id)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"payload":room})
	return
}
