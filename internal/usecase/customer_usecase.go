package usecase

import (
	"context"
	"github.com/bookingtogo/internal/entity"
	"github.com/bookingtogo/internal/model"
	"github.com/bookingtogo/internal/model/converter"
	"github.com/bookingtogo/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
		return nil, fiber.ErrBadRequest
	}

	Customer := &entity.Customer{
		Title:    request.Title,
		Content:  request.Content, // harusnya open langsung default
		Category: request.Category,
		Status:   request.Status,
	}

	if err := c.CustomerRepository.Create(tx, Customer); err != nil {
		c.Log.WithError(err).Error("error creating Customer")
		return nil, fiber.ErrInternalServerError

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating Customer")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerToResponse(Customer), nil

}

func (c *CustomerUseCase) Update(ctx context.Context,
	request *model.UpdateCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	Customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, Customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, fiber.ErrNotFound
	}
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Input yang dimasukan ada kesalahan")
	}

	if request.Title != "" {
		Customer.Title = request.Title
	}

	if request.Content != "" {
		Customer.Content = request.Content
	}

	if request.Category != "" {
		Customer.Category = request.Category
	}

	if request.Status != "" {
		Customer.Status = request.Status
	}

	if err := c.CustomerRepository.Update(tx, Customer); err != nil {
		c.Log.WithError(err).Error("error Update Customer")
		return nil, fiber.ErrInternalServerError

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error Update Customer")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerToResponse(Customer), nil

}

func (c *CustomerUseCase) Get(ctx context.Context, request *model.GetCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	Customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, Customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerToResponse(Customer), nil
}

func (c *CustomerUseCase) Delete(ctx context.Context, request *model.DeleteCustomerRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}

	Customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, Customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return fiber.ErrNotFound
	}

	if err := c.CustomerRepository.Delete(tx, Customer); err != nil {
		c.Log.WithError(err).Error("error deleting Customer")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting Customer")
		return http.ErrInternalServerError
	}

	return nil
}

func (c *CustomerUseCase) Search(ctx context.Context, request *model.SearchCustomerRequest) ([]model.CustomerResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	contacts, total, err := c.CustomerRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting contacts")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting contacts")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.CustomerResponse, len(contacts))
	for i, contact := range contacts {
		responses[i] = *converter.CustomerToResponse(&contact)
	}

	return responses, total, nil
}
func (c *CustomerUseCase) FindAll(ctx context.Context, request *model.AllCustomerRequest) ([]model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	Customer, err := c.CustomerRepository.FindAll(tx)
	if err != nil {
		c.Log.Warnf("Failed find all site : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.CustomerResponse, len(Customer))
	for i, localSite := range Customer {
		responses[i] = *converter.CustomerToResponse(&localSite)
	}

	return responses, nil
}
