package model

type Article struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title" validate:"required,min=20"`
	Content     string `json:"content" validate:"required,min=200"`
	Category    string `json:"category" validate:"required,min=3"`
	CreatedDate string `json:"created_date,omitempty"`
	UpdatedDate string `json:"updated_date,omitempty"`
	Status      string `json:"status" validate:"required,oneof=Publish Draft Thrash"`
}
