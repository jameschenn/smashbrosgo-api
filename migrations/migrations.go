package main

import (
	"github.com/jameschenn/smashbrosgo-api/models"
	"github.com/jameschenn/smashbrosgo-api/utils"
)

func init() {
	utils.LoadEnvVariables()
	utils.ConnectToDB()
}

func main() {
	utils.DB.AutoMigrate(&models.Character{})
}
