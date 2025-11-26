package usecase

import (
	"context"
	"github.com/bookingtogo/internal/helper"
	"github.com/bookingtogo/internal/model"
	"github.com/bookingtogo/internal/model/converter"
	"github.com/bookingtogo/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type NationalityUseCase struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	Validate              *validator.Validate
	NationalityRepository *repository.NationalityRepository
}

func NewNationalityUseCase(db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	NationalityRepository *repository.NationalityRepository) *NationalityUseCase {
	return &NationalityUseCase{
		DB:                    db,
		Log:                   logger,
		Validate:              validate,
		NationalityRepository: NationalityRepository,
	}
}

func (c *NationalityUseCase) FindAll(ctx context.Context, request *model.GetNationalityRequest) ([]model.NationalityResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, helper.NewBadRequest("Input ada yang salah", err)
	}

	Nationality, err := c.NationalityRepository.FindAll(tx)
	if err != nil {
		c.Log.Warnf("Failed find all user : %+v", err)
		return nil, helper.NewInternal("Internal service error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, helper.NewInternal("Internal service error")
	}
	// Cek kalau data kosong
	if len(Nationality) == 0 {
		tx.Rollback()
		return nil, helper.NewNotFound("nationality data is currently empty.") // atau helper lain sesuai kebutuhan
	}

	responses := make([]model.NationalityResponse, len(Nationality))
	for i, Nationality := range Nationality {
		responses[i] = *converter.NationalityToResponse(&Nationality)
	}

	return responses, nil
}
