package model

type PostResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	Status      string `json:"status"`
}

type CreatePostRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft trash"`
}

type UpdatePostRequest struct {
	ID       string `json:"-" validate:"required"`
	Title    string `json:"title" validate:"min=20"`
	Content  string `json:"content" validate:"min=200"`
	Category string `json:"category"`
	Status   string `json:"status" validate:"required,oneof=publish draft trash"`
}

type GetPostRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeletePostRequest struct {
	ID string `json:"id" validate:"required"`
}
type SearchPostRequest struct {
	ID       string `json:"-" validate:""`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
	Page     int    `json:"page" validate:"min=1"`
	Size     int    `json:"size" validate:"min=1,max=100"`
}

type AllPostRequest struct {
	ID       string `json:"-" validate:""`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}
