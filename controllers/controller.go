package controllers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	UploadFilePath := filepath.Join(folderpath, File.Filename)
	if err := c.SaveUploadedFile(File, UploadFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't save uploaded file", "error": err.Error()})
		return
	}
	DataChan := make(chan []string)
	defer close(DataChan)
	go func(filePath string) {
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

		defer file.Close()

		DataChan <- FileDatas

	}(folderpath)

	fileData := <-DataChan
	Text := make(map[string]int)
	for _, eachText := range fileData {
		words := strings.Fields(eachText)
		for _, eachWord := range words {
			Text[eachWord]++
		}

	}
	c.JSON(http.StatusOK, gin.H{"set": Text})
}
