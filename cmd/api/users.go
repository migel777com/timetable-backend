package main

import (
	"gin-api-template/internal/data"
	"gin-api-template/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)
//refactor error structs

// UserCreate
// @Summary      User creation
// @Description  Send all data about user, except photo_id. Plz exclude id, createdTime and photo_id fields in request
// @Accept       json
// @Produce      json
// @Param        user body data.RegisterUser true "User data"
// @Success      200  {object}  int
// @Failure      422  {object}  data.RegisterUserValidationError
// @Failure      406  {object}  SimpleError
// @Failure      500  {object}  SimpleError
// @Router       /users/register/new [post]
func (app *application) UserCreate(c *gin.Context) {

	//getting user data from request
	var input data.RegisterUser

	if err := c.BindJSON(&input); err != nil {
		app.serverErrorResponse(err, c)
	}

	//checking if user with this email already exists
	possibleUser, err := app.models.Users.GetByEmail(input.Email)
	if possibleUser!=nil{
		app.NotAcceptable(err, c)
		return
	}

	//moving input data to user struct to access validating functionality
	user := &data.User{
		Email:       input.Email,
		Password:    input.Password,
		Phone:       input.Phone,
		Name:        input.Name,
		Surname:     input.Surname,
		Country:     input.Country,
		City:        input.City,
		Address:     input.Address,
	}

	//validating it
	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		app.FailedValidationResponse(v.Errors, c)
		return
	}

	//replace password with hashed password
	newHashedPass, err := HashPassword(user.Password)
	if err!=nil{
		app.serverErrorResponse(err, c)
	}

	user.Password = newHashedPass

	//insert
	insert, err := app.models.Users.Insert(user.Email, user.Password, user.Phone, user.Name, user.Surname, user.Country, user.City, user.Address)
	if err != nil {
		app.serverErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"payload":insert})
	return
}

// EmailAuth
// @Summary      Auth via Email
// @Description  Send email and pass to get JWT
// @Accept       json
// @Produce      json
// @Param        user body data.LoginUserViaEmail true "Login data"
// @Success      200  {object}  data.SimplePayload
// @Failure      422  {object}  data.LoginUserViaEmailValidationError
// @Failure      401  {object}  SimpleError
// @Failure      500  {object}  SimpleError
// @Router       /users/login/email [post]
func (app *application) EmailAuth(c *gin.Context){
	var input data.LoginUserViaEmail

	if err := c.BindJSON(&input); err != nil {
		app.serverErrorResponse(err, c)
	}

	//moving input data to user struct to access validating functionality
	user := &data.User{
		Email:       input.Email,
		Password:    input.Password,
	}

	//validating it
	v := validator.New()
	if data.ValidateLogin(v, user); !v.Valid() {
		app.FailedValidationResponse(v.Errors, c)
		return
	}

	possibleUser, err := app.models.Users.GetByEmail(user.Email)
	if err!=nil{
		if err.Error()=="sql: no rows in result set"{
			app.InvalidCredentials(err, c)
			return
		}
		app.serverErrorResponse(err, c)
		return
	}

	hashStatus, err := CompareHash(user.Password, possibleUser.Password)
	if err!=nil{
		app.serverErrorResponse(err, c)
		return
	}
	if hashStatus{
		newJWT, err := GenerateJWT(app.config.Jwtkey, user)
		if err!=nil{
			app.serverErrorResponse(err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{"payload":newJWT})
		return
	}else{
		app.InvalidCredentials(err, c)
		return
	}
}

// EmailAuth
// @Summary      Get user data with get request
// @Description  use JWT token to identify users email and respond with data related to  this user
// @Produce      json
// @Param        user body data.LoginUserViaEmail true "Login data"
// @Success      200  {object}  data.SimplePayload
// @Failure      401  {object}  SimpleError
// @Failure      500  {object}  SimpleError
// @Router       /users/userdata [get]
func (app *application) UserData(c *gin.Context){
	//get user email from headers
	claims, _ := extractClaims(app.config.Jwtkey, c.Request.Header["Akis-Jwt-Token"][0])
	userEmail := claims["email"].(string)
	user, err := app.models.Users.GetByEmail(userEmail)
	if err!=nil{

		if err.Error()=="sql: no rows in result set"{
			app.InvalidCredentials(err, c)
			return
		}
		app.serverErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"payload":user})
	return
}


