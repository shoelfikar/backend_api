package repository

import "github.com/shoelfikar/golang_backend_api/model"

type UserRepository interface {
	Create(u *model.User) *model.User
	GetUserById(userId int) *model.User
	GetUserByEmail(login *model.LoginRequest) *model.User
}