package water

type Water struct {
	ID             int
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
	Water             Water
}
