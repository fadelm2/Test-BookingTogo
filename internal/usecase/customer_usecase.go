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
	"strconv"
	"time"
)

type CustomerUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	CustomerRepository   *repository.CustomerRepository
	FamilyListRepository *repository.FamilyListRepository
}

func NewCustomerUseCase(db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	CustomerRepository *repository.CustomerRepository,
	FamilyListRepository *repository.FamilyListRepository) *CustomerUseCase {
	return &CustomerUseCase{
		DB:                   db,
		Log:                  logger,
		Validate:             validate,
		CustomerRepository:   CustomerRepository,
		FamilyListRepository: FamilyListRepository,
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

	parsedDob, err := time.Parse("2006-01-02", request.Dob)
	if err != nil {
		c.Log.Error(err)
	}

	Customer := &entity.Customer{
		Name:          request.Name,
		DOB:           parsedDob, // harusnya open langsung default
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

func (c *CustomerUseCase) FindAll(ctx context.Context, request *model.AllCustomerRequest) ([]model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, helper.NewBadRequest("Input ada yang salah", err)
	}

	Customers, err := c.CustomerRepository.FindAll(tx)
	if err != nil {
		c.Log.Warnf("Failed find all user : %+v", err)
		return nil, helper.NewInternal("Internal service error")
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, helper.NewInternal("Internal service error")
	}
	// Cek kalau data kosong
	if len(Customers) == 0 {
		tx.Rollback()
		return nil, helper.NewNotFound("nationality data is currently empty.") // atau helper lain sesuai kebutuhan
	}

	responses := make([]model.CustomerResponse, len(Customers))
	for i, customer := range Customers {
		responses[i] = *converter.CustomerToResponse(&customer)
	}

	return responses, nil
}

func (c *CustomerUseCase) CreateWithFamily(
	ctx context.Context,
	request *model.CreateCustomerWithFamilyRequest,
) (*model.CustomerWithFamilyResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// 1. VALIDATION
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, helper.NewBadRequest("Input is incorrect", err)
	}

	// 2. PARSE DOB
	parsedDob, err := time.Parse("2006-01-02", request.Dob)
	if err != nil {
		c.Log.WithError(err).Error("invalid date format")
		return nil, helper.NewBadRequest("DOB must be in format YYYY-MM-DD", err)
	}

	// 3. CREATE CUSTOMER ENTITY
	customer := &entity.Customer{
		Name:          request.Name,
		DOB:           parsedDob,
		NationalityId: request.NationalityID,
		Email:         request.Email,
		Phone:         request.PhoneNumber,
	}

	// 4. INSERT CUSTOMER
	if err := c.CustomerRepository.Create(tx, customer); err != nil {
		c.Log.WithError(err).Error("failed to create customer")
		return nil, helper.NewInternal("Internal service error")
	}

	// 5. INSERT FAMILY LIST (LOOP)
	var familyEntities []entity.FamilyList
	c.Log.Debugf("berikut data request :", request.FamilyRequest)
	for _, famReq := range request.FamilyRequest {

		// Optional: validasi masing² family
		if err := c.Validate.Struct(famReq); err != nil {
			return nil, helper.NewBadRequest("Family input invalid", err)
		}

		if err != nil {
			return nil, helper.NewBadRequest("Family DOB format invalid", err)
		}

		family := entity.FamilyList{
			CustomerID: customer.ID,
			Name:       famReq.Name,
			Relation:   famReq.Relation,
			Dob:        famReq.Dob,
		}

		if err := c.FamilyListRepository.Create(tx, &family); err != nil {
			c.Log.WithError(err).Error("failed to create family")
			return nil, helper.NewInternal("Internal service error")
		}

		familyEntities = append(familyEntities, family)
	}

	// 6. COMMIT TRANSACTION
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed commit")
		return nil, helper.NewInternal("Internal service error")
	}

	// 7. RETURN RESPONSE (gabungan customer + family)
	return converter.CustomerWithFamilyToResponse(customer, familyEntities), nil
}

func (c *CustomerUseCase) UpdateWithFamily(
	ctx context.Context,
	request *model.UpdateCustomerWithFamilyRequest,
) (*model.CustomerWithFamilyResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// 1. VALIDATION
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, helper.NewBadRequest("Input is incorrect", err)
	}

	// 2. CHECK CUSTOMER EXISTS
	Customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, Customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return nil, helper.NewNotFound("customer id not found")
	}

	if request.Name != nil {
		Customer.Name = *request.Name
	}

	if request.Dob != nil {
		parsedDob, err := time.Parse("2006-01-02", *request.Dob)
		if err != nil {
			return nil, helper.NewBadRequest("DOB must be YYYY-MM-DD", err)
		}
		Customer.DOB = parsedDob
	}

	if request.Email != nil {
		Customer.Email = *request.Email
	}

	if request.PhoneNumber != nil {
		Customer.Phone = *request.PhoneNumber
	}

	if err := c.CustomerRepository.Update(tx, Customer); err != nil {
		c.Log.WithError(err).Error("failed to update customer")
		return nil, helper.NewInternal("Internal service error")
	}

	// ====================================
	// 5. HANDLE FAMILY LIST UPDATE
	// ====================================

	// Ambil data family yang sudah ada

	existingFamilies, err := c.FamilyListRepository.FindAllFamily(tx, strconv.Itoa(Customer.ID))
	if err != nil {
		// Error dari DB → return 500
		c.Log.WithError(err).Error("error fetching family list")
		return nil, helper.NewInternal("failed to get family data")
	}

	existingMap := map[int]entity.FamilyList{}
	for _, f := range existingFamilies {
		existingMap[f.ID] = f
	}

	var updatedFamilies []entity.FamilyList

	for _, famReq := range request.FamilyRequest {

		// Validasi input tiap family
		if err := c.Validate.Struct(famReq); err != nil {
			return nil, helper.NewBadRequest("Family input invalid", err)
		}

		// CASE 1: Update family (ada ID)
		if famReq.ID != 0 {
			famEntity, exists := existingMap[famReq.ID]
			if !exists {
				return nil, helper.NewBadRequest("Family ID not found", nil)
			}

			// Update value
			famEntity.Name = famReq.Name
			famEntity.Relation = famReq.Relation
			famEntity.Dob = famReq.Dob

			if err := c.FamilyListRepository.Update(tx, &famEntity); err != nil {
				return nil, helper.NewInternal("Failed updating family")
			}

			updatedFamilies = append(updatedFamilies, famEntity)
			delete(existingMap, famReq.ID) // tandai sebagai tidak dihapus
			continue
		}

		// CASE 2: Insert family baru (tanpa ID)
		newFam := entity.FamilyList{
			CustomerID: Customer.ID,
			Name:       famReq.Name,
			Relation:   famReq.Relation,
			Dob:        famReq.Dob,
		}

		if err := c.FamilyListRepository.Create(tx, &newFam); err != nil {
			return nil, helper.NewInternal("Failed creating family")
		}

		updatedFamilies = append(updatedFamilies, newFam)
	}

	// CASE 3: Delete family yang tidak ada di permintaan (opsional)
	for _, famToDelete := range existingMap {
		if err := c.FamilyListRepository.DeleteByID(tx, famToDelete.ID); err != nil {
			return nil, helper.NewInternal("Failed deleting family")
		}
	}

	// ====================================
	// 6. COMMIT
	// ====================================
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed commit")
		return nil, helper.NewInternal("Internal service error")
	}

	// 7. RETURN RESPONSE
	return converter.CustomerWithFamilyToResponse(Customer, updatedFamilies), nil
}

func (c *CustomerUseCase) GetCustomerWithFamily(
	ctx context.Context,
	request *model.GetFamilyListRequest,
) (*model.CustomerWithFamilyResponse, error) {

	tx := c.DB.WithContext(ctx)

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, helper.NewBadRequest("Input is incorrect", err)
	}
	// 1. GET CUSTOMER
	Customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, Customer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return nil, helper.NewNotFound("customer ID not found")
	}

	// 2. GET FAMILY LIST
	families, err := c.FamilyListRepository.FindAllFamily(tx, request.ID)
	if err != nil {
		c.Log.WithError(err).Error("failed finding family list")
		return nil, helper.NewInternal("Internal server error")
	}

	// jika tidak ada family → pakai slice kosong (lebih aman untuk FE)
	if len(families) == 0 {
		families = []entity.FamilyList{}
	}

	// 3. RETURN RESULT
	return converter.CustomerWithFamilyToResponse(Customer, families), nil
}
