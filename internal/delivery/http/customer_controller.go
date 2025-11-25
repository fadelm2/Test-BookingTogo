package http

import (
	"github.com/bookingtogo/internal/model"
	"github.com/bookingtogo/internal/usecase"
	"math"

	"github.com/sirupsen/logrus"
)

type CustomerController struct {
	UseCase *usecase.CustomerUseCase
	Log     *logrus.Logger
}

func NewCustomerController(useCase *usecase.CustomerUseCase, log *logrus.Logger) *CustomerController {
	return &CustomerController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *CustomerController) Create(ctx *.Ctx) error {

	request := new(model.CreateCustomerRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating Customer")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Data: response})
}

func (c *CustomerController) List(ctx *fiber.Ctx) error {

	request := &model.SearchCustomerRequest{
		Title:    ctx.Query("title", ""),
		Content:  ctx.Query("name", ""),
		Category: ctx.Query("category", ""),
		Status:   ctx.Query("status", ""),
		Page:     ctx.QueryInt("page", 1),
		Size:     ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching Customer")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.CustomerResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *CustomerController) Get(ctx *fiber.Ctx) error {

	request := &model.GetCustomerRequest{
		ID: ctx.Params("CustomerId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting Customer")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Data: response})
}

func (c *CustomerController) Update(ctx *fiber.Ctx) error {

	request := new(model.UpdateCustomerRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("CustomerId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating Customer")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Data: response})
}

func (c *CustomerController) Delete(ctx *fiber.Ctx) error {
	CustomerId := ctx.Params("CustomerId")

	request := &model.DeleteCustomerRequest{
		ID: CustomerId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("error deleting Customer")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

func (c *CustomerController) FindAllCustomer(ctx *fiber.Ctx) error {

	request := &model.AllCustomerRequest{
		ID: ctx.Query("id", ""),
	}

	responses, err := c.UseCase.FindAll(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get all Article")
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.CustomerResponse]{
		Data: responses,
	})
}
