package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"fmt"
	"time"
	"io"
)

type BlogInput struct{
	Author string `json:"author"`
	Title string `json:"title"`
	Comments []struct{
		Message string `json:"message"`
	} `json:"comments"`
}



func main() {
	r:= gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

	// your custom format
	return fmt.Sprintf("[%s] \"%s %s %d\n",
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.StatusCode,
		)
	}))
	r.Use(gin.Recovery())

	// r := gin.Default()

	r.POST("/blogs", func(c *gin.Context) {
		f, _ := os.Create("author.log")
		gin.DefaultWriter = io.MultiWriter(f)

		// authorFile,err := os.OpenFile("author.log",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		// titleFile,err := os.OpenFile("title.log",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		// defer authorFile.Close()
		// defer titleFile.Close()

		// loggerAuthor := log.New(authorFile, "", log.LstdFlags)
		// loggerTitle := log.New(titleFile, "", log.LstdFlags)
		
		var dataInput BlogInput

		if err := c.ShouldBindJSON(&dataInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		// loggerAuthor.Printf("{\"author\": \"%s\"}",dataInput.Author)
		// loggerTitle.Printf("{\"title\": \"%s\"}",dataInput.Title)

		c.JSON(http.StatusCreated, gin.H{"author": dataInput.Author,"title": dataInput.Title,"comments": dataInput.Comments})
	})

	r.Run()
}