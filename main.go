package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	utils "github.com/jameschenn/smashbrosgo-api/utils"
)

func init() {
	utils.LoadEnvVariables()
	utils.ConnectToDB()
}

func main() {
	fmt.Println("Hello Gopher James")
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
