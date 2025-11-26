package model

type CustomerResponse struct {
	ID            int    `json:"id"`
	NationalityID int    `json:"nationality_id"`
	Name          string `json:"name"`
	Dob           string `json:"dob"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
}

type CustomerWithFamilyResponse struct {
	ID            int                   `json:"id"`
	NationalityID int                   `json:"nationality_id"`
	Name          string                `json:"name"`
	Dob           string                `json:"dob"`
	PhoneNumber   string                `json:"phone_number"`
	Email         string                `json:"email"`
	Family        []*FamilyListResponse `json:"families"`
}
type CreateCustomerRequest struct {
	NationalityID int    `json:"nationality_id" validate:"required"`
	Name          string `json:"name" validate:"required,min=3"`
	Dob           string `json:"dob" validate:"required"`
	PhoneNumber   string `json:"phone_number" validate:"min=13"`
	Email         string `json:"email" validate:"email"`
}
type CreateCustomerWithFamilyRequest struct {
	NationalityID int                       `json:"nationality_id" validate:"required"`
	Name          string                    `json:"name" validate:"required,min=3"`
	Dob           string                    `json:"dob" validate:"required"`
	PhoneNumber   string                    `json:"phone_number" validate:"min=13"`
	Email         string                    `json:"email" validate:"email"`
	FamilyRequest []CreateFamilyListRequest `json:"families"`
}

type UpdateCustomerRequest struct {
	ID            string  `json:"-" validate:"required"`
	NationalityID int     `json:"nationality_id"`
	Name          string  `json:"name" validate:"min=3"`
	Dob           *string `json:"dob"`
	PhoneNumber   string  `json:"phone_number" validate:"min=14"`
	Email         string  `json:"email" validate:"email"`
}

type GetCustomerRequest struct {
	ID int `json:"-" validate:"required"`
}

type DeleteCustomerRequest struct {
	ID int `json:"-" validate:"required"`
}

type SearchCustomerRequest struct {
	ID            string `json:"-"`
	NationalityID int    `json:"nationality_id"`
	Name          string `json:"name"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
	Page          int    `json:"page" validate:"min=1"`
	Size          int    `json:"size" validate:"min=1,max=100"`
}

type AllCustomerRequest struct {
	ID            string `json:"-"`
	NationalityID int    `json:"nationality_id"`
	Name          string `json:"name"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
}

type UpdateCustomerWithFamilyRequest struct {
	ID            string                `json:"-"`
	Name          *string               `json:"name"`
	Dob           *string               `json:"dob"`
	Email         *string               `json:"email"`
	PhoneNumber   *string               `json:"phone_number"`
	NationalityID *int                  `json:"nationality_id"`
	FamilyRequest []UpdateFamilyRequest `json:"families"`
}

type UpdateFamilyRequest struct {
	ID       int    `json:"id"` // nil = create new
	Name     string `json:"name"`
	Relation string `json:"relation"`
	Dob      string `json:"dob"`
}
