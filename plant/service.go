package plant

type Service interface {
	CreatePlant(ID int) (Plant, error)
	GetUserPlant(ID int) (Plant, error)
	UpdatePlant(plan Plant) (Plant, error)
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

func (s *service) UpdatePlant(plant Plant) (Plant, error) {

	updatedPlant, err := s.repository.Update(plant)
	if err != nil {
		return updatedPlant, err
	}
	return updatedPlant, nil
}
