package water

type Service interface {
	CreateUserWater(userID int, input BuyWaterEnergy) (UserWater, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUserWater(userID int, input BuyWaterEnergy) (UserWater, error) {
	userWater := UserWater{}
	userWater.UserID = userID
	userWater.WaterEnergyTypeID = input.WaterEnergyTypeID
	userWater.Qty = input.Qty
	userWater.Water, _ = s.repository.GetByID(userWater.WaterEnergyTypeID)
	newUserWater, err := s.repository.SaveUserWater(userWater)
	if err != nil {
		return newUserWater, err
	}
	return newUserWater, nil
}
