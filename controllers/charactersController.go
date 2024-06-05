package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jameschenn/smashbrosgo-api/models"
	"github.com/jameschenn/smashbrosgo-api/utils"
)

func CharacterCreate(c *gin.Context) {

	// Get data off request body
	var payload struct {
		Name            string
		Series          string
		FirstAppearance string
		Description     string
		Tier            string
		Weight          float64
		Speed           float64
		NeutralB        string
		UpB             string
		SideB           string
		DownB           string
		Final_Smash     string
		Image_URL       string
	}

	c.Bind(&payload)

	// Create Character
	character := models.Character{
		Name:            payload.Name,
		Series:          payload.Series,
		FirstAppearance: payload.FirstAppearance,
		Description:     payload.Description,
		Tier:            payload.Tier,
		Weight:          payload.Weight,
		Speed:           payload.Speed,
		NeutralB:        payload.NeutralB,
		UpB:             payload.UpB,
		SideB:           payload.SideB,
		DownB:           payload.DownB,
		Final_Smash:     payload.Final_Smash,
		Image_URL:       payload.Image_URL,
	}

	result := utils.DB.Create(&character)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return Character
	c.JSON(200, gin.H{
		"character": character,
	})
}

func GetAllCharacters(c *gin.Context) {

	//Get all characters
	var characters []models.Character
	utils.DB.Find(&characters)

	// Return array of all characters
	c.JSON(200, gin.H{
		"characters": characters,
	})
}

func GetCharacterByID(c *gin.Context) {

	//Get id from request body
	id := c.Param("id")

	//Get all characters
	var character models.Character
	utils.DB.First(&character, id)

	// Returns characters
	c.JSON(200, gin.H{
		"character": character,
	})
}

func UpdateCharacterByID(c *gin.Context) {

	// Get id from request body
	id := c.Param("id")

	// Get data off request body
	var payload struct {
		Name            string
		Series          string
		FirstAppearance string
		Description     string
		Tier            string
		Weight          float64
		Speed           float64
		NeutralB        string
		UpB             string
		SideB           string
		DownB           string
		Final_Smash     string
		Image_URL       string
	}

	c.Bind(&payload)

	//Find the character we're updating
	var character models.Character
	utils.DB.First(&character, id)

	//Update the character
	utils.DB.Model(&character).Updates(models.Character{
		Name:            payload.Name,
		Series:          payload.Series,
		FirstAppearance: payload.FirstAppearance,
		Description:     payload.Description,
		Tier:            payload.Tier,
		Weight:          payload.Weight,
		Speed:           payload.Speed,
		NeutralB:        payload.NeutralB,
		UpB:             payload.UpB,
		SideB:           payload.SideB,
		DownB:           payload.DownB,
		Final_Smash:     payload.Final_Smash,
		Image_URL:       payload.Image_URL,
	})

	//Return character
	c.JSON(200, gin.H{
		"character": character,
	})
}

func DeleteCharacterByID(c *gin.Context) {

	// Get id from request body
	id := c.Param("id")

	//Delete the post
	utils.DB.Delete(&models.Character{}, id)

	//Respond
	c.Status(200)
}
