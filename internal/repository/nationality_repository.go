package repository

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/sirupsen/logrus"
)

type NationalityRepository struct {
	Repository[entity.Nationality]
	Log *logrus.Logger
}

func NewNationalityRepository(log *logrus.Logger) *NationalityRepository {
	return &NationalityRepository{
		Log: log,
	}
}
