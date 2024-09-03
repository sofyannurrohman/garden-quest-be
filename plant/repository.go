package plant

import "gorm.io/gorm"

type Repository interface {
	SavePlant(plant Plant) (Plant, error)
	FetchUserPlant(ID int) (Plant, error)
	FindByID(ID int) (Plant, error)
	Update(plan Plant) (Plant, error)
	FindAllPlantType() ([]PlantType, error)
	FindPlantTypeByID(ID int) (PlantType, error)
	SaveUserPlant(userPlant Plant) (Plant, error)
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
func (r *repository) FetchUserPlant(ID int) (Plant, error) {
	var plant Plant
	err := r.db.Preload("PlantType").Where("user_id = ?", ID).Find(&plant).Error

	if err != nil {
		return plant, err
	}
	return plant, nil
}
func (r *repository) FindByID(ID int) (Plant, error) {
	var plant Plant
	err := r.db.Where("id = ?", ID).Find(&plant).Error
	if err != nil {
		return plant, err
	}
	return plant, nil
}

func (r *repository) Update(plant Plant) (Plant, error) {
	err := r.db.Save(&plant).Error
	if err != nil {
		return plant, err
	}
	return plant, nil
}

func (r *repository) FindPlantTypeByID(ID int) (PlantType, error) {
	var plantType PlantType
	err := r.db.Where("id = ?", ID).Find(&plantType).Error
	if err != nil {
		return plantType, err
	}
	return plantType, nil
}

func (r *repository) FindAllPlantType() ([]PlantType, error) {
	var plantType []PlantType
	err := r.db.Find(&plantType).Error
	if err != nil {
		return plantType, err
	}
	return plantType, nil
}

func (r *repository) SaveUserPlant(userPlant Plant) (Plant, error) {
	err := r.db.Create(&userPlant).Error
	if err != nil {
		return userPlant, err
	}
	return userPlant, nil
}
