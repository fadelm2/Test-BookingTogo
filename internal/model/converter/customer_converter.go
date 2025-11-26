package converter

import (
	"github.com/bookingtogo/internal/entity"
	"github.com/bookingtogo/internal/model"
)

func CustomerToResponse(c *entity.Customer) *model.CustomerResponse {

	return &model.CustomerResponse{
		ID:            c.ID,
		NationalityID: c.NationalityId,
		Name:          c.Name,
		Dob:           c.DOB.Format("2006-01-02"),
		PhoneNumber:   c.Phone,
		Email:         c.Email,
	}
}

func FamilyListToResponseList(list []entity.FamilyList) []*model.FamilyListResponse {
	var result []*model.FamilyListResponse

	for _, f := range list {
		result = append(result, FamilyListToResponse(&f))
	}

	return result
}

func CustomerToResponseWithFamily(c *entity.Customer) *model.CustomerWithFamilyResponse {

	return &model.CustomerWithFamilyResponse{
		ID:            c.ID,
		NationalityID: c.NationalityId,
		Name:          c.Name,
		Dob:           c.DOB.Format("2006-01-02"),
		PhoneNumber:   c.Phone,
		Email:         c.Email,
		Family:        FamilyListToResponseList(c.Family),
	}
}

func CustomerWithFamilyToResponse(
	customer *entity.Customer,
	family []entity.FamilyList,
) *model.CustomerWithFamilyResponse {

	response := &model.CustomerWithFamilyResponse{
		ID:            customer.ID,
		NationalityID: customer.NationalityId,
		Name:          customer.Name,
		Dob:           customer.DOB.Format("2006-01-02"),
		PhoneNumber:   customer.Phone,
		Email:         customer.Email,
	}

	// convert family list
	var familyResponse []*model.FamilyListResponse

	for _, f := range family {
		familyResponse = append(familyResponse, &model.FamilyListResponse{
			ID:         f.ID,
			CustomerID: f.CustomerID,
			Relation:   f.Relation,
			Name:       f.Name,
			Dob:        f.Dob,
		})
	}

	response.Family = familyResponse

	return response
}
