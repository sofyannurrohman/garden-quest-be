package user

type UserFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	ImageURL    string `json:"image_url"`
	Coins       int    `json:"coins"`
	WaterEnergy int    `json:"water_energy"`
}

func FormatJSONUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Coins:       user.Coins,
		WaterEnergy: user.WaterEnergy,
		Token:       token,
		ImageURL:    user.AvatarFileName,
	}
	return formatter
}
