package repository

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/sirupsen/logrus"
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
