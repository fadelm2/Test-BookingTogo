package model

type FamilyListResponse struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	Relation   string `json:"relation"`
	Name       string `json:"name"`
	Dob        string `json:"dob"`
}

type CreateFamilyListRequest struct {
	CustomerID int    `json:"-" validate:""`
	Relation   string `json:"relation" validate:""`
	Name       string `json:"name" validate:""`
	Dob        string `json:"dob" validate:""`
}

type UpdateFamilyListRequest struct {
	ID         string `json:"-" validate:"required"`
	CustomerID int    `json:"customer_id"`
	Relation   string `json:"relation" validate:"min=3"`
	Name       string `json:"name" validate:"min=3"`
	Dob        string `json:"dob"`
}

type GetFamilyListRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteFamilyListRequest struct {
	ID string `json:"id" validate:"required"`
}

type SearchFamilyListRequest struct {
	ID         string `json:"-"`
	CustomerID int    `json:"customer_id"`
	Relation   string `json:"relation"`
	Name       string `json:"name"`
	Page       int    `json:"page" validate:"min=1"`
	Size       int    `json:"size" validate:"min=1,max=100"`
}

type AllFamilyListRequest struct {
	ID         string `json:"-"`
	CustomerID int    `json:"customer_id"`
	Relation   string `json:"relation"`
	Name       string `json:"name"`
}
