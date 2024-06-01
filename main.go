package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/jameschenn/smashbrosgo-api/controllers"
	utils "github.com/jameschenn/smashbrosgo-api/utils"
)

func init() {
	utils.LoadEnvVariables()
	utils.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/characters", controllers.CharacterCreate)
	router.Run() // listen and serve on 0.0.0.0:8080
}
