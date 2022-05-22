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


