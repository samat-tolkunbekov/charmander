package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPages(c *gin.Context) {
	fmt.Println("Hello stranger")

	c.JSON(http.StatusOK, gin.H{
		"message": "Which page you want to see?",
	})

	return
}

func GetUserData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "You are good my friend",
	})
}
