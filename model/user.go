package model

type User struct {
	Id        int     `json:"id"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	Status    int     `json:"status"`
	Token     *string `json:"token,omitempty"`
	CreatedAt string  `json:"created_at"`
	CreatedBy string  `json:"created_by"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
