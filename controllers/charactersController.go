package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jameschenn/smashbrosgo-api/models"
	"github.com/jameschenn/smashbrosgo-api/utils"
)

func CharacterCreate(c *gin.Context) {

	// Get data off request body

	// Create Character
	character := models.Character{
		Name:            "James Chen",
		Series:          "James Chen",
		FirstAppearance: "1992",
		Description:     "Mario is an all-around fighter who uses his wide variety of techniques to respond to any situation. In Super Smash Bros. Ultimate, he shows up in his wedding tux and builder outfit, and Cappy even makes an appearance!",
		Tier:            "A",
		Weight:          160,
		Speed:           3.33,
		NeutralB:        "Strong Communication!!",
		UpB:             "Rising Coding Skills!!",
		SideB:           "Collaborative Spirit!!",
		DownB:           "Critical Thinking Charge!!",
		Final_Smash:     "Got the job!!",
		Image_URL:       "https://avatars.githubusercontent.com/u/73676915?v=4",
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
