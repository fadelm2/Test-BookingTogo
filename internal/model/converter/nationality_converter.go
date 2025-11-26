package converter

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/bookingtogo/internal/model"
)

func NationalityToResponse(r *entity.Nationality) *model.NationalityResponse {

	return &model.NationalityResponse{
		ID:   r.NationalityId,
		Name: r.Name,
		Code: r.Code,
	}
}
