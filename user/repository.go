package user

import (
	"garden-quest/plant"
	"garden-quest/water"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
	FindAll() ([]User, error)
	FindAllPlants(userID int) ([]plant.Plant, error)
	FindAllWaters(userID int) ([]water.UserWater, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
func (r *repository) FindAllPlants(userID int) ([]plant.Plant, error) {
	var plants []plant.Plant
	err := r.db.Where("user_id = ?", userID).Find(&plants).Error
	if err != nil {
		return []plant.Plant{}, err
	}
	return plants, nil
}

func (r *repository) FindAllWaters(userID int) ([]water.UserWater, error) {
	var waters []water.UserWater
	err := r.db.Where("user_id = ?", userID).Find(&waters).Error
	if err != nil {
		return []water.UserWater{}, err
	}
	return waters, nil
}
