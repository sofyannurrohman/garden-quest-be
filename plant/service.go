package plant

type Service interface {
	CreatePlant(ID int) (Plant, error)
	GetUserPlant(ID int) (Plant, error)
	GetPlantTypeByID(ID int) (PlantType, error)
	UpdatePlant(plan Plant) (Plant, error)
	CreateUserPlant(userID int, input BuyPlant) (Plant, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreatePlant(ID int) (Plant, error) {
	plant := Plant{}
	plant.UserID = ID
	plant.PlantTypeID = 1
	plant.WateringCount = 0

	newPlant, err := s.repository.SavePlant(plant)
	if err != nil {
		return newPlant, err
	}
	return newPlant, nil
}

func (s *service) GetUserPlant(ID int) (Plant, error) {

	plant, err := s.repository.FetchUserPlant(ID)
	if err != nil {
		return plant, err
	}
	return plant, nil
}

func (s *service) GetPlantTypeByID(ID int) (PlantType, error) {

	plantType, err := s.repository.FindPlantTypeByID(ID)
	if err != nil {
		return plantType, err
	}
	return plantType, nil
}

func (s *service) UpdatePlant(plant Plant) (Plant, error) {

	updatedPlant, err := s.repository.Update(plant)
	if err != nil {
		return updatedPlant, err
	}
	return updatedPlant, nil
}

func (s *service) CreateUserPlant(userID int, input BuyPlant) (Plant, error) {
	userPlant := Plant{}
	userPlant.PlantTypeID = input.PlantTypeID
	userPlant.UserID = userID
	userPlant.WateringCount = 0
	userPlant.PlantType, _ = s.GetPlantTypeByID(userPlant.PlantTypeID)
	newUserPlant, err := s.repository.SavePlant(userPlant)
	if err != nil {
		return newUserPlant, err
	}
	return newUserPlant, nil
}
