package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	targetDir = "D:\\java笔记\\system-file"
)

func saveFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving file")
		panic(err)
	}
	err = os.MkdirAll(targetDir, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Faliled to create")
		panic(err)
	}

	fileName := "test.txt"
	filePath := targetDir + "\\" + fileName
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save uploaded file: %v", err)
		panic(err)
	}
}
