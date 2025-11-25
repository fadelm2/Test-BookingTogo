package model

type NationalityResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreateNationalityRequest struct {
	Name string `json:"name" validate:"required,min=3"`
	Code string `json:"code" validate:"required,len=2"`
}

type UpdateNationalityRequest struct {
	ID   string `json:"-" validate:"required"`
	Name string `json:"name" validate:"min=3"`
	Code string `json:"code" validate:"len=2"`
}

type GetNationalityRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteNationalityRequest struct {
	ID string `json:"id" validate:"required"`
}

type SearchNationalityRequest struct {
	ID   string `json:"-"`
	Name string `json:"name"`
	Code string `json:"code"`
	Page int    `json:"page" validate:"min=1"`
	Size int    `json:"size" validate:"min=1,max=100"`
}

type AllNationalityRequest struct {
	ID   string `json:"-"`
	Name string `json:"name"`
	Code string `json:"code"`
}
