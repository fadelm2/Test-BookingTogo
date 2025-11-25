package http

import (
	"github.com/bookingtogo/internal/model"
	"github.com/bookingtogo/internal/usecase"
	"math"

	"github.com/sirupsen/logrus"
)

type PostController struct {
	UseCase *usecase.PostUseCase
	Log     *logrus.Logger
}

func NewPostController(useCase *usecase.PostUseCase, log *logrus.Logger) *PostController {
	return &PostController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *PostController) Create(ctx *fiber.Ctx) error {

	request := new(model.CreatePostRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating Post")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.PostResponse]{Data: response})
}

func (c *PostController) List(ctx *fiber.Ctx) error {

	request := &model.SearchPostRequest{
		Title:    ctx.Query("title", ""),
		Content:  ctx.Query("name", ""),
		Category: ctx.Query("category", ""),
		Status:   ctx.Query("status", ""),
		Page:     ctx.QueryInt("page", 1),
		Size:     ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching Post")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.PostResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *PostController) Get(ctx *fiber.Ctx) error {

	request := &model.GetPostRequest{
		ID: ctx.Params("PostId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting Post")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.PostResponse]{Data: response})
}

func (c *PostController) Update(ctx *fiber.Ctx) error {

	request := new(model.UpdatePostRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("PostId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating Post")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.PostResponse]{Data: response})
}

func (c *PostController) Delete(ctx *fiber.Ctx) error {
	PostId := ctx.Params("PostId")

	request := &model.DeletePostRequest{
		ID: PostId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("error deleting Post")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

func (c *PostController) FindAllPost(ctx *fiber.Ctx) error {

	request := &model.AllPostRequest{
		ID: ctx.Query("id", ""),
	}

	responses, err := c.UseCase.FindAll(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get all Article")
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.PostResponse]{
		Data: responses,
	})
}
