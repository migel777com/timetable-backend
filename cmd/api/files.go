package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

// FileUpload
// @Summary      Uploading files
// @Description  Send as binary file and write name of the file in the URL
// @Accept       json
// @Produce      json
// @Param        file body data.File true "File data"
// @Success      200  {object}  int
// @Failure      500  {object}  SimpleError
// @Router       /files/upload/{fileName} [post]
func (app *application) FileUpload(c *gin.Context) {

	name := c.Param("fileName")
	body := c.Request.Body
	photo, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}

	path := "uploads/"

	tempFile, err := ioutil.TempFile(path, "*-" + name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	_, err = tempFile.Write(photo)
	if err != nil {
		fmt.Println(err)
		return
	}

	ext := filepath.Ext(tempFile.Name())


	insert, err := app.models.Files.Insert(strings.TrimLeft(tempFile.Name(), "uploads\\"), ext, tempFile.Name())
	if err != nil {
		app.serverErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"payload":insert})
	return

}
