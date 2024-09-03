package user

import (
	"garden-quest/plant"
	"garden-quest/water"
	"time"
)

type User struct {
	ID             int `gorm:"primary_key"`
	Name           string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Coins          int
	WaterEnergy    int
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Inventory struct {
	Plants []plant.Plant
	Waters []water.UserWater
}
