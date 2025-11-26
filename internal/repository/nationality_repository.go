package repository

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (r *NationalityRepository) FindAll(db *gorm.DB) ([]entity.Nationality, error) {
	var nationalities []entity.Nationality

	if err := db.Find(&nationalities).Error; err != nil {
		return nil, err
	}

	return nationalities, nil
}
