package service

import (
	"errors"

	"github.com/shoelfikar/golang_backend_api/helper"
	"github.com/shoelfikar/golang_backend_api/model"
	"github.com/shoelfikar/golang_backend_api/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	UserRepository repository.UserRepository
}

type UserService interface {
	CreateUser(usr *model.User) *model.User
	GetUserById(userId int) *model.User
	GenerateLoginToken(user *model.LoginRequest) *model.User
}

func NewUserService(user repository.UserRepository) UserService {
	return &userService{
		UserRepository: user,
	}
}

func (sv *userService) CreateUser(usr *model.User) *model.User {

	password := []byte(usr.Password)
	// Generate a bcrypt hash from the password
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	helper.PanicIfError(err)

	usr.Password = string(hashedPassword)

	result := sv.UserRepository.Create(usr)

	return result
}

func (sv *userService) GetUserById(userId int) *model.User {
	return sv.UserRepository.GetUserById(userId)
}

func (sv *userService) GenerateLoginToken(user *model.LoginRequest) *model.User{
	existUser := sv.UserRepository.GetUserByEmail(user)

	storedHashPass := []byte(existUser.Password)
	userPassword := []byte(user.Password)

	errDecrypt := bcrypt.CompareHashAndPassword(storedHashPass, userPassword)
	
	helper.PanicIfError(errDecrypt)

	token := helper.GetViperEnvVariable("API_KEY")

	if token == "" {
		tokenErr := errors.New("token not found")
		helper.PanicIfError(tokenErr)
	}

	existUser.Token = &token

	return existUser
}
