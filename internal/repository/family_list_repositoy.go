package repository

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FamilyListRepository struct {
	Repository[entity.FamilyList]
	Log *logrus.Logger
}

func NewFamilyListRepository(log *logrus.Logger) *FamilyListRepository {
	return &FamilyListRepository{
		Log: log,
	}
}

func (r *FamilyListRepository) FindAllFamily(db *gorm.DB, customerID string) ([]entity.FamilyList, error) {
	var families []entity.FamilyList

	if err := db.Where("cst_id = ?", customerID).Find(&families).Error; err != nil {
		return nil, err
	}

	return families, nil
}
