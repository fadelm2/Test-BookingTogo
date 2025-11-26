package repository

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	Repository[entity.Customer]
	Log *logrus.Logger
}

func NewCustomerRepository(log *logrus.Logger) *CustomerRepository {
	return &CustomerRepository{
		Log: log,
	}
}
func (r *CustomerRepository) FindAll(db *gorm.DB) ([]entity.Customer, error) {
	var customer []entity.Customer

	if err := db.Find(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}
