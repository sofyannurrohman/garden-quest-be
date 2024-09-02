package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserInput struct {
	WaterEnergy int `json:"water_energy" `
}
type AddEnergy struct {
	AddWaterEnergy int `json:"add_water_energy"`
}
type FormCreateUserInput struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Error    error
}
type FormUpdateUserInput struct {
	ID    int
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required,email"`
	Error error
}
