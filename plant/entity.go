package plant

type Plant struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	PlantTypeID   int       `json:"plant_type_id"`
	WateringCount int       `json:"watering_count"`
	PlantType     PlantType `gorm:"foreignKey:PlantTypeID"`
}

type PlantType struct {
	ID           int    `gorm:"primary_key"`
	Name         string `json:"name"`
	Goal         int    `json:"goal"`
	CoinProduced int    `json:"coin_produced"`
	Price        int    `json:"price"`
}

type UserPlant struct {
	ID          int
	UserID      int
	PlantTypeID int
	PlantType   PlantType
}
