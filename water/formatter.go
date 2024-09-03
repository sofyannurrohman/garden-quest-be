package water

type WaterFormatter struct {
	ID                int `json:"id"`
	WaterEnergyTypeID int `json:"water_energy_type_id"`
	Qty               int `json:"qty"`
	UserID            int `json:"user_id"`
}

func FormatJSONWater(userwater UserWater) WaterFormatter {
	formatter := WaterFormatter{
		ID:                userwater.ID,
		WaterEnergyTypeID: userwater.WaterEnergyTypeID,
		Qty:               userwater.Qty,
		UserID:            userwater.UserID,
	}
	return formatter
}
