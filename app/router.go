package app

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/shoelfikar/golang_backend_api/controller"
	"github.com/shoelfikar/golang_backend_api/exception"
	"github.com/shoelfikar/golang_backend_api/helper"
	"github.com/shoelfikar/golang_backend_api/middleware"
	"github.com/shoelfikar/golang_backend_api/repository"
	"github.com/shoelfikar/golang_backend_api/service"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	validate := validator.New()

	router.Use(exception.Recovery)
	router.Use(middleware.CORS)
	router.Use(middleware.LoggingMiddleware)
	
	// web := router.PathPrefix("/web").Subrouter()
	nonAuth := router.PathPrefix("/api/v1").Subrouter()
	auth := router.PathPrefix("/api/v1").Subrouter()

	//auth middleware
	auth.Use(middleware.IsAuthorized)
	
	//user controller
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	//article controller
	articleRepo := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepo, validate)
	articleController := controller.NewArticleController(articleService)

	//web controller
	webController := controller.NewWebController()
	
	// non auth
	nonAuth.HandleFunc("/user", userController.CreateUser).Methods("POST")
	nonAuth.HandleFunc("/auth/login", userController.Login).Methods("POST")

	//auth route
	auth.HandleFunc("/article", articleController.CreateArticle).Methods("POST")
	auth.HandleFunc("/article/{id:[0-9]}", articleController.GetArticleById).Methods("GET")
	auth.HandleFunc("/article/{id:[0-9]}", articleController.UpdateArticleById).Methods("PUT")
	auth.HandleFunc("/article/{id:[0-9]}", articleController.DeleteArticleById).Methods("DELETE")
	auth.HandleFunc("/article/{limit:[0-9]*}/{offset:[0-9]*}", articleController.GetArticles).Methods("GET")

	//web route
	router.HandleFunc("/login", webController.LoginWebController)
	router.HandleFunc("/", webController.DashboardWebController)
	
	helper.DefaultLoggingInfo("server running on port " + helper.GetViperEnvVariable("PORT"))
	
	return router

}