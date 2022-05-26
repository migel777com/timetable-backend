package main

import (
	"gin-api-template/internal/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) OfficeAuth(c *gin.Context){
	var input data.User

	if err := c.BindJSON(&input); err != nil {
		app.serverErrorResponse(err, c)
	}

	user := &data.User{
		Email:       input.Email,
		Organization:input.Organization,
	}

	newJWT, err := GenerateJWT(app.config.Jwtkey, user)
	if err!=nil{
		app.serverErrorResponse(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"payload":newJWT})
	return
}

func (app *application) isAdmin(c *gin.Context) {
	id := c.Param("id")

	admin, err := app.models.Users.GetAdmin(id)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.NotFoundResponse(err, c)
			return
		} else {
			app.serverErrorResponse(err, c)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"payload":admin})
	return
}


