package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"log"
	"encoding/json"
	"time"
)

type BlogInput struct{
	Author string `json:"author"`
	Title string `json:"title"`
	Comments []struct{
		Message string `json:"message"`
	} `json:"comments"`
}



func main() {
	r := gin.Default()

	r.POST("/blogs", func(c *gin.Context) {

		authorFile,err := os.OpenFile("author.log",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		titleFile,err := os.OpenFile("title.log",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		commentsFile,err := os.OpenFile("comments.log",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open errors log file")
		}

		defer authorFile.Close()
		defer titleFile.Close()
		defer commentsFile.Close()

		loggerAuthor := log.New(authorFile, "", 0)
		loggerTitle := log.New(titleFile, "", 0)
		loggerComments := log.New(commentsFile, "", 0)
		
		var dataInput BlogInput

		if err := c.ShouldBindJSON(&dataInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		author,_:=json.Marshal(dataInput.Author)
		title,_:=json.Marshal(dataInput.Title)
		comments,_:=json.Marshal(dataInput.Comments)

		loggerAuthor.Printf("[%s] %d {\"author\": %s}",
			time.Now().Format(time.RFC3339),
			http.StatusCreated,
			string(author))

		loggerTitle.Printf("[%s] %d {\"title\": %s}",
			time.Now().Format(time.RFC3339),
			http.StatusCreated,
			string(title))

		loggerComments.Printf("[%s] %d {\"comments\": %s}",
			time.Now().Format(time.RFC3339),
			http.StatusCreated,
			string(comments))

		c.JSON(http.StatusCreated, gin.H{"author": dataInput.Author,"title": dataInput.Title,"comments": dataInput.Comments})
	})

	r.Run()
}