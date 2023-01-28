package main

import (
	"charmander/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("Hello world")

	r := gin.Default()

	server := r.Group("/pages")
	{
		server.GET("/branch", func(c *gin.Context) {
			internal.GetPages(c)
		})

		server.GET("/get-user-data", func(c *gin.Context) {
			internal.GetUserData(c)
		})

		server.POST("/get-data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data new returned",
				"body":    "Here you go",
			})
		})
	}

	err := r.Run(":9898")

	if err != nil {
		return
	}
}
