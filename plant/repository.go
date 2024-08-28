package plant

import "gorm.io/gorm"

type Repository interface {
	SavePlant(plant Plant) (Plant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SavePlant(plant Plant) (Plant, error) {
	err := r.db.Create(&plant).Error
	if err != nil {
		return plant, err
	}
	return plant, nil
}
