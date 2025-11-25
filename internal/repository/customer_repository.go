package repository

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/sirupsen/logrus"
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
