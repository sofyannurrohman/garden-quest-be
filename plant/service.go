package plant

type Service interface {
	CreatePlant(ID int) (Plant, error)
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
