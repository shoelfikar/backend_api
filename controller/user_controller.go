package controller

import (
	"net/http"

	"github.com/shoelfikar/golang_backend_api/helper"
	"github.com/shoelfikar/golang_backend_api/model"
	"github.com/shoelfikar/golang_backend_api/service"
)

type userController struct{
	userService service.UserService
}

type UserController interface{
	CreateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

func NewUserController(usr service.UserService) UserController {
	return &userController{
		userService: usr,
	}
}

func (u *userController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user *model.User

	helper.ReadFromRequestBody(r, &user)


	resp := u.userService.CreateUser(user)

	response := model.ApiResponse{
		Code: http.StatusCreated,
		Status: "success",
		Message: "OK",
		Data: resp,
	}


	helper.WriteToResponseBody(w, response)
	
}

func (u *userController) Login(w http.ResponseWriter, r *http.Request) {
	var userLogin model.LoginRequest

	helper.ReadFromRequestBody(r, &userLogin)

	user := u.userService.GenerateLoginToken(&userLogin)

	response := model.ApiResponse{
		Code: http.StatusOK,
		Status: "success",
		Message: "OK",
		Data: user,
	}


	helper.WriteToResponseBody(w, response)
}
