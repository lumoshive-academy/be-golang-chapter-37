package repository

import (
	"golang-chapter-37/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll() ([]model.Customer, error)
	Save(customer *model.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindAll() ([]model.Customer, error) {
	var customers []model.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *customerRepository) Save(customer *model.Customer) error {
	return r.db.Create(customer).Error
}
