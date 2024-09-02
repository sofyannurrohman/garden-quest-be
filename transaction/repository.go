package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetAllTransaction() ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) GetAllTransaction() ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
