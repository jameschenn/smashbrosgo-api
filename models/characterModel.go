package models

import (
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
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
