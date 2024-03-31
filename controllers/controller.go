package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FileReader(c *gin.Context) {
	File, err := c.FormFile("file")
	fmt.Println("call received")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "something wrong in the request", "error": err.Error()})
		return
	}
	folderpath := "/home/lenovo/FileProcessorr/files"
	UploadFilePath := folderpath + File.Filename
	if err := c.SaveUploadedFile(File, UploadFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't save uplaoded file", "error": err.Error()})
		return
	}
	dat, err := os.ReadFile(UploadFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't read file", "error": err.Error()})
		return
	}
	fmt.Println("Data", dat)
}
