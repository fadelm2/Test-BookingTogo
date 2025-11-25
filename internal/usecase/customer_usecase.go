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
	"time"
)

type CustomerUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	CustomerRepository *repository.CustomerRepository
}

func NewCustomerUseCase(db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	CustomerRepository *repository.CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		CustomerRepository: CustomerRepository,
	}
}

func (c *CustomerUseCase) Create(ctx context.Context,
	request *model.CreateCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, helper.NewBadRequest("Input is incorrect", err)
	}

	Customer := &entity.Customer{
		Name:          request.Name,
		DOB:           request.Dob, // harusnya open langsung default
		NationalityId: request.NationalityID,
		Email:         request.Email,
		Phone:         request.PhoneNumber,
	}

	if err := c.CustomerRepository.Create(tx, Customer); err != nil {
		c.Log.WithError(err).Error("error creating Customer")
		return nil, helper.NewInternal("Internal service error")

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating Customer")
		return nil, helper.NewInternal("Internal service error")
	}

	return converter.CustomerToResponse(Customer), nil

}

func (c *CustomerUseCase) Update(ctx context.Context,
	request *model.UpdateCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	Customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, Customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return nil, helper.NewNotFound("Customer Id Not Fund")
	}
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, helper.NewBadRequest("Input is incorrect", err)
	}

	if request.Name != "" {
		Customer.Name = request.Name
	}

	if request.Dob != nil {
		dob, err := time.Parse("2006-01-02", *request.Dob)
		if err != nil {
			helper.NewBadRequest("Input is incorrect", err)
		}
		Customer.DOB = dob
	}
	if request.PhoneNumber != "" {
		Customer.Phone = request.PhoneNumber
	}

	if request.Email != "" {
		Customer.Email = request.Email
	}

	if err := c.CustomerRepository.Update(tx, Customer); err != nil {
		c.Log.WithError(err).Error("error Update Customer")
		return nil, helper.NewInternal("Internal service error")

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error Update Customer")
		return nil, helper.NewInternal("Internal service error")
	}

	return converter.CustomerToResponse(Customer), nil

}

func (c *CustomerUseCase) Get(ctx context.Context, request *model.GetCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, helper.NewBadRequest("Input ada yang salah", err)
	}

	Customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, Customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return nil, helper.NewNotFound("customer id not found")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return nil, helper.NewInternal("Internal service error")
	}

	return converter.CustomerToResponse(Customer), nil
}

func (c *CustomerUseCase) Delete(ctx context.Context, request *model.DeleteCustomerRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return helper.NewBadRequest("Input ada yang salah", err)
	}

	Customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, Customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return helper.NewNotFound("Customer Id Not Fund")
	}

	if err := c.CustomerRepository.Delete(tx, Customer); err != nil {
		c.Log.WithError(err).Error("error deleting Customer")
		return helper.NewInternal("Internal service error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting Customer")
		return helper.NewInternal("Internal service error")
	}

	return nil
}

//func (c *CustomerUseCase) Search(ctx context.Context, request *model.SearchCustomerRequest) ([]model.CustomerResponse, int64, error) {
//	tx := c.DB.WithContext(ctx).Begin()
//	defer tx.Rollback()
//
//	if err := c.Validate.Struct(request); err != nil {
//		c.Log.WithError(err).Error("error validating request body")
//		return nil, 0, helper.NewBadRequest("Input ada yang salah", err)
//	}
//
//	contacts, total, err := c.CustomerRepository.Search(tx, request)
//	if err != nil {
//		c.Log.WithError(err).Error("error getting contacts")
//		return nil, 0, helper.NewInternal("Internal service error")
//	}
//
//	if err := tx.Commit().Error; err != nil {
//		c.Log.WithError(err).Error("error getting contacts")
//		return nil, 0, helper.NewInternal("Internal service error")
//	}
//
//	responses := make([]model.CustomerResponse, len(contacts))
//	for i, contact := range contacts {
//		responses[i] = *converter.CustomerToResponse(&contact)
//	}
//
//	return responses, total, nil
//}
//func (c *CustomerUseCase) FindAll(ctx context.Context, request *model.AllCustomerRequest) ([]model.CustomerResponse, error) {
//	tx := c.DB.WithContext(ctx).Begin()
//	defer tx.Rollback()
//
//	if err := c.Validate.Struct(request); err != nil {
//		c.Log.WithError(err).Error("error validating request body")
//		return nil, helper.NewBadRequest("Input ada yang salah", err)
//	}
//
//	Customer, err := c.CustomerRepository.FindAll(tx)
//	if err != nil {
//		c.Log.Warnf("Failed find all site : %+v", err)
//		return nil, helper.NewInternal("Internal service error")
//	}
//
//	if err := tx.Commit().Error; err != nil {
//		c.Log.Warnf("Failed commit transaction : %+v", err)
//		return nil, helper.NewInternal("Internal service error")
//	}
//
//	responses := make([]model.CustomerResponse, len(Customer))
//	for i, localSite := range Customer {
//		responses[i] = *converter.CustomerToResponse(&localSite)
//	}
//
//	return responses, nil
//}
