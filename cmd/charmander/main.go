package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello world")

	r := gin.Default()

	//r.LoadHTMLGlob("../../web/static/*.tmpl")

	server := r.Group("/server")
	{
		server.POST("/get-data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data returned",
				"body":    "Here you go",
			})

			//c.HTML(http.StatusOK, "index.tmpl", gin.H{
			//	"title": "Hello world!",
			//})
		})
	}

	err := r.Run(":" + os.Getenv("APP_PORT"))

	if err != nil {
		return
	}
}
