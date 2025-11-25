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

type PostUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	PostRepository *repository.PostsRepository
}

func NewPostUseCase(db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PostRepository *repository.PostsRepository) *PostUseCase {
	return &PostUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		PostRepository: PostRepository,
	}
}

func (c *PostUseCase) Create(ctx context.Context,
	request *model.CreatePostRequest) (*model.PostResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	Post := &entity.Post{
		Title:    request.Title,
		Content:  request.Content, // harusnya open langsung default
		Category: request.Category,
		Status:   request.Status,
	}

	if err := c.PostRepository.Create(tx, Post); err != nil {
		c.Log.WithError(err).Error("error creating Post")
		return nil, fiber.ErrInternalServerError

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating Post")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PostToResponse(Post), nil

}

func (c *PostUseCase) Update(ctx context.Context,
	request *model.UpdatePostRequest) (*model.PostResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	post := new(entity.Post)
	if err := c.PostRepository.FindById(tx, post, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, fiber.ErrNotFound
	}
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Input yang dimasukan ada kesalahan")
	}

	if request.Title != "" {
		post.Title = request.Title
	}

	if request.Content != "" {
		post.Content = request.Content
	}

	if request.Category != "" {
		post.Category = request.Category
	}

	if request.Status != "" {
		post.Status = request.Status
	}

	if err := c.PostRepository.Update(tx, post); err != nil {
		c.Log.WithError(err).Error("error Update Post")
		return nil, fiber.ErrInternalServerError

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error Update Post")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PostToResponse(post), nil

}

func (c *PostUseCase) Get(ctx context.Context, request *model.GetPostRequest) (*model.PostResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	post := new(entity.Post)
	if err := c.PostRepository.FindById(tx, post, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Post")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting Post")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PostToResponse(post), nil
}

func (c *PostUseCase) Delete(ctx context.Context, request *model.DeletePostRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}

	post := new(entity.Post)
	if err := c.PostRepository.FindById(tx, post, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting Post")
		return fiber.ErrNotFound
	}

	if err := c.PostRepository.Delete(tx, post); err != nil {
		c.Log.WithError(err).Error("error deleting Post")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting Post")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *PostUseCase) Search(ctx context.Context, request *model.SearchPostRequest) ([]model.PostResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	contacts, total, err := c.PostRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting contacts")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting contacts")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.PostResponse, len(contacts))
	for i, contact := range contacts {
		responses[i] = *converter.PostToResponse(&contact)
	}

	return responses, total, nil
}
func (c *PostUseCase) FindAll(ctx context.Context, request *model.AllPostRequest) ([]model.PostResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	Post, err := c.PostRepository.FindAll(tx)
	if err != nil {
		c.Log.Warnf("Failed find all site : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.PostResponse, len(Post))
	for i, localSite := range Post {
		responses[i] = *converter.PostToResponse(&localSite)
	}

	return responses, nil
}
