package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		if err := c.SaveUploadedFile(file, "images/" + file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success!!"})
}
