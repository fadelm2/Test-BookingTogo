package converter

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/bookingtogo/internal/model"
)

func PostToResponse(post *entity.Post) *model.PostResponse {
	return &model.PostResponse{
		ID:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		Category:    post.Category,
		Status:      post.Status,
		CreatedDate: post.CreatedDate.Format("2006-01-02 15:04:05"),
		UpdatedDate: post.CreatedDate.Format("2006-01-02 15:04:05"),
	}
}
