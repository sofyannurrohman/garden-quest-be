package user

import (
	"errors"
	"garden-quest/plant"
	"garden-quest/water"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
	GetAllUser() ([]User, error)
	AddWater(userID int, input UpdateUserInput) (User, error)
	AddEnergy(userID int, input AddEnergy) (User, error)
	BuyWaterEnergy(userID int, input water.BuyWaterEnergy) (water.UserWater, error)
	BuyPlantType(userID int, input plant.BuyPlant) (plant.Plant, error)
	GetInventory(userID int) (Inventory, error)
}
type service struct {
	repository      Repository
	plantService    plant.Service
	waterRepository water.Repository
	waterService    water.Service
}

func NewService(repository Repository, plantService plant.Service, waterRepository water.Repository, waterService water.Service) *service {
	return &service{repository, plantService, waterRepository, waterService}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Coins = 0
	user.WaterEnergy = 3
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	_, err = s.plantService.CreatePlant(newUser.ID)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
func (s *service) LoginUser(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}
func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	//getuser by id
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	//update attribute avatar filename
	user.AvatarFileName = fileLocation
	//save update to db
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil

}
func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}
	return user, nil
}
func (s *service) GetAllUser() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}
func (s *service) AddWater(userID int, input UpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}
	plant, err := s.plantService.GetUserPlant(userID)
	if err != nil {
		return user, err
	}
	user.WaterEnergy -= input.WaterEnergy
	newWaterCount := plant.WateringCount + input.WaterEnergy

	plant.WateringCount = newWaterCount

	_, err = s.plantService.UpdatePlant(plant)
	if err != nil {
		return user, err
	}
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func (s *service) AddEnergy(userID int, input AddEnergy) (User, error) {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return user, err
	}
	user.WaterEnergy += input.AddWaterEnergy
	newUser, err := s.repository.Update(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) BuyWaterEnergy(userID int, input water.BuyWaterEnergy) (water.UserWater, error) {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return water.UserWater{}, err
	}
	waterType, err := s.waterRepository.GetByID(input.WaterEnergyTypeID)
	if err != nil {
		return water.UserWater{}, err
	}
	price := waterType.Price * input.Qty
	if user.Coins < price {
		return water.UserWater{}, err
	}
	user.Coins -= price
	newWater, err := s.waterService.CreateUserWater(userID, input)
	if err != nil {
		return water.UserWater{}, err
	}
	_, err = s.repository.Update(user)
	if err != nil {
		return water.UserWater{}, err
	}
	return newWater, nil
}

func (s *service) BuyPlantType(userID int, input plant.BuyPlant) (plant.Plant, error) {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return plant.Plant{}, err
	}
	planType, err := s.plantService.GetPlantTypeByID(input.PlantTypeID)
	if err != nil {
		return plant.Plant{}, err
	}
	if user.Coins < planType.Price {
		return plant.Plant{}, err
	}
	user.Coins -= planType.Price
	newPlant, err := s.plantService.CreateUserPlant(userID, input)
	if err != nil {
		return newPlant, err
	}
	_, err = s.repository.Update(user)
	if err != nil {
		return newPlant, err
	}
	return newPlant, err
}

func (s *service) GetInventory(userID int) (Inventory, error) {

	userPlants, err := s.repository.FindAllPlants(userID)
	if err != nil {
		return Inventory{}, err
	}
	waters, err := s.repository.FindAllWaters(userID)
	if err != nil {
		return Inventory{}, err
	}
	inventory := Inventory{}
	inventory.Plants = userPlants
	inventory.Waters = waters
	return inventory, nil
}
