package converter

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/bookingtogo/internal/model"
)

func FamilyListToResponse(f *entity.FamilyList) *model.FamilyListResponse {

	return &model.FamilyListResponse{
		ID:         f.ID,
		CustomerID: f.CustomerID,
		Relation:   f.Relation,
		Name:       f.Name,
		Dob:        f.Dob.Format("2006-01-02"),
	}
}
