package water

type WaterType struct {
	ID             int `gorm:"primary_key"`
	Name           string
	Description    string
	AddWaterEnergy int
	Price          int
}

type UserWater struct {
	ID                int
	WaterEnergyTypeID int
	Qty               int
	UserID            int
	Water             WaterType `gorm:"foreignKey:WaterEnergyTypeID"`
}
