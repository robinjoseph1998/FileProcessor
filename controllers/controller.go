package controllers

import (
	"bufio"
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

	file, err := os.Open(UploadFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't open the file", "error": err.Error()})
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var FileDatas []string

	for scanner.Scan() {
		FileDatas = append(FileDatas, scanner.Text())
	}

	file.Close()

	for _, eachLine := range FileDatas {
		fmt.Println(eachLine)
	}
}
