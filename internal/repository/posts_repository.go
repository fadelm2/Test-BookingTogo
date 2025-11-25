package repository

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/bookingtogo/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostsRepository struct {
	Repository[entity.Post]
	Log *logrus.Logger
}

func NewPostsRepository(log *logrus.Logger) *PostsRepository {
	return &PostsRepository{
		Log: log,
	}
}

func (r *PostsRepository) Search(db *gorm.DB, request *model.SearchPostRequest) ([]entity.Post, int64, error) {
	var Postss []entity.Post
	if err := db.Scopes(r.FilterPosts(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&Postss).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Post{}).Scopes(r.FilterPosts(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return Postss, total, nil
}

func (r *PostsRepository) FilterPosts(request *model.SearchPostRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {

		if PostsID := request.ID; PostsID != "" {
			PostsID = "%" + PostsID + "%"
			tx = tx.Where("id LIKE ? ", PostsID)
		}

		if Title := request.Title; Title != "" {
			Title = "%" + Title + "%"
			tx = tx.Where("Title LIKE ?", Title)
		}

		if Content := request.Content; Content != "" {
			Content = "%" + Content + "%"
			tx = tx.Where("Content LIKE ?", Content)
		}
		if Content := request.Content; Content != "" {
			Content = "%" + Content + "%"
			tx = tx.Where("Content LIKE ?", Content)
		}

		if Category := request.Category; Category != "" {
			Category = "%" + Category + "%"
			tx = tx.Where("Category LIKE ?", Category)
		}
		if Status := request.Status; Status != "" {
			Status = "%" + Status + "%"
			tx = tx.Where("Status LIKE ?", Status)
		}

		return tx
	}
}

func (r *PostsRepository) UpdatePosts(db *gorm.DB, entity *entity.Post) error {
	return db.Save(&entity).Error
}

func (r *PostsRepository) FindAll(db *gorm.DB) ([]entity.Post, error) {
	var Post []entity.Post
	if err := db.Scopes().Find(&Post).Error; err != nil {
		return nil, err
	}

	return Post, nil
}
