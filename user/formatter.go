package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
	
}

func FormatJSONUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Token:    token,
		ImageURL: user.AvatarFileName,
	}
	return formatter
}
