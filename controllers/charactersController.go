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
