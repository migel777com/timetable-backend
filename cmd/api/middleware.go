package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (app *application) auth() gin.HandlerFunc {
	return func (c *gin.Context) {
		if c.Request.Header["Gao-Jwt-Token"] != nil {

			token, err := jwt.Parse(c.Request.Header["Gao-Jwt-Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("SigningMethodHMAC checking error")
				}
				return app.config.Jwtkey, nil
			})

			if err != nil {
				app.serverErrorResponse(err, c)
				return
			}

			if token.Valid {
				claims, _ := extractClaims(app.config.Jwtkey, c.Request.Header["Gao-Jwt-Token"][0])
				organization := claims["organization"]

				if organization == "Astana IT University" {
					c.Next()
				} else {
					app.InvalidCredentials(nil, c)
					return
				}
			}
		} else {

			app.NotAuthorized(nil, c)
			c.Abort()
			return
		}
	}
}