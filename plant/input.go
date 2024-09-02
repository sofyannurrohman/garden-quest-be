package plant

type UpdatePlantInput struct {
	UserID        int `json:"user_id"`
	PlantTypeID   int `json:"plant_type_id"`
	WateringCount int `json:"watering_count"`
}
type BuyPlant struct {
	PlantTypeID int `json:"plant_type_id"`
}
