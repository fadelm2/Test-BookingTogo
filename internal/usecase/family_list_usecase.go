package usecase

import (
	"context"
	"github.com/bookingtogo/internal/entity"
	"github.com/bookingtogo/internal/helper"
	"github.com/bookingtogo/internal/model"
	"github.com/bookingtogo/internal/model/converter"
	"github.com/bookingtogo/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FamilyListUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	FamilyListRepository *repository.FamilyListRepository
	CustomerRepository   *repository.CustomerRepository
}

func NewFamilyListUseCase(db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	FamilyListRepository *repository.FamilyListRepository,
	CustomerRepository *repository.CustomerRepository) *FamilyListUseCase {
	return &FamilyListUseCase{
		DB:                   db,
		Log:                  logger,
		Validate:             validate,
		FamilyListRepository: FamilyListRepository,
		CustomerRepository:   CustomerRepository,
	}
}

func (c *FamilyListUseCase) Create(ctx context.Context,
	request *model.CreateFamilyListRequest) (*model.FamilyListResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, helper.NewBadRequest("Input is incorrect", err)
	}

	Customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, Customer, request.CustomerID); err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return nil, helper.NewNotFound("customer id not found")
	}

	FamilyList := &entity.FamilyList{
		Name:       request.Name,
		Dob:        request.Dob, // harusnya open langsung default
		Relation:   request.Relation,
		CustomerID: request.CustomerID,
	}

	if err := c.FamilyListRepository.Create(tx, FamilyList); err != nil {
		c.Log.WithError(err).Error("error creating FamilyList")
		return nil, helper.NewInternal("Internal service error")

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating FamilyList")
		return nil, helper.NewInternal("Internal service error")
	}

	return converter.FamilyListToResponse(FamilyList), nil

}

func (c *FamilyListUseCase) Delete(ctx context.Context, request *model.DeleteFamilyListRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return helper.NewBadRequest("Input ada yang salah", err)
	}

	FamilyList := new(entity.FamilyList)
	if err := c.FamilyListRepository.FindById(tx, FamilyList, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting FamilyList")
		return helper.NewNotFound("FamilyList Id Not Fund")
	}

	if err := c.FamilyListRepository.Delete(tx, FamilyList); err != nil {
		c.Log.WithError(err).Error("error deleting FamilyList")
		return helper.NewInternal("Internal service error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting FamilyList")
		return helper.NewInternal("Internal service error")
	}

	return nil
}

func (c *FamilyListUseCase) FindAll(ctx context.Context, request model.GetFamilyListRequest) ([]model.FamilyListResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, helper.NewBadRequest("Input errors", err)
	}

	Customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, Customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return nil, helper.NewNotFound("customer id not found")
	}

	FamilyList, err := c.FamilyListRepository.FindAllFamily(tx, request.ID)
	if err != nil {
		c.Log.Warnf("Failed find all user : %+v", err)
		return nil, helper.NewInternal("Internal service error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, helper.NewInternal("Internal service error")
	}
	// Cek kalau data kosong
	if len(FamilyList) == 0 {
		tx.Rollback()
		return nil, helper.NewNotFound("Family data is currently empty.") // atau helper lain sesuai kebutuhan
	}

	responses := make([]model.FamilyListResponse, len(FamilyList))
	for i, familyList := range FamilyList {
		responses[i] = *converter.FamilyListToResponse(&familyList)
	}

	return responses, nil
}
