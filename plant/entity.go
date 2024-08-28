package plant

type Plant struct {
	ID            int `json:"id"`
	UserID        int `json:"user_id"`
	PlantTypeID   int `json:"plant_type_id"`
	WateringCount int `json:"watering_count"`
}
