package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jameschenn/smashbrosgo-api/controllers"
	utils "github.com/jameschenn/smashbrosgo-api/utils"
)

func init() {
	utils.LoadEnvVariables()
	utils.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.GET("/characters", controllers.GetAllCharacters)
	router.GET("/characters/:id", controllers.GetCharacterByID)
	router.POST("/characters", controllers.CharacterCreate)
	router.PUT("/characters/:id", controllers.UpdateCharacterByID)
	router.DELETE("characters/:id", controllers.DeleteCharacterByID)
	router.Run() // listen and serve on 0.0.0.0:8080
}
