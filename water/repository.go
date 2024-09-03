package water

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]WaterType, error)
	GetByID(ID int) (WaterType, error)
	SaveUserWater(userWater UserWater) (UserWater, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]WaterType, error) {
	var water []WaterType
	err := r.db.Find(&water).Error
	if err != nil {
		return nil, err
	}
	return water, nil
}

func (r *repository) GetByID(ID int) (WaterType, error) {
	var water WaterType
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
