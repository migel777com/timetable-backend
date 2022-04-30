package main

import (
	data "gin-api-template/internal/data"
	"github.com/gin-gonic/gin"
	"net/http"
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


func (app *application) Home(c *gin.Context) {

}
