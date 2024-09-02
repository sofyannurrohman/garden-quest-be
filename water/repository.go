package water

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Water, error)
	GetByID(ID int) (Water, error)
	SaveUserWater(userWater UserWater) (UserWater, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Water, error) {
	var water []Water
	err := r.db.Find(&water).Error
	if err != nil {
		return nil, err
	}
	return water, nil
}

func (r *repository) GetByID(ID int) (Water, error) {
	var water Water
	err := r.db.Where("id=?", ID).Find(&water).Error
	if err != nil {
		return water, err
	}
	return water, nil
}

func (r *repository) SaveUserWater(userWater UserWater) (UserWater, error) {
	err := r.db.Create(&userWater).Error
	if err != nil {
		return userWater, err
	}
	return userWater, nil
}
