package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shoelfikar/golang_backend_api/helper"
	"github.com/shoelfikar/golang_backend_api/model"
	"github.com/shoelfikar/golang_backend_api/service"
)

type articleController struct {
	articleService service.ArticleService
}

type ArticleController interface{
	CreateArticle(w http.ResponseWriter, r *http.Request)
	GetArticles(w http.ResponseWriter, r *http.Request)
	GetArticleById(w http.ResponseWriter, r *http.Request)
	UpdateArticleById(w http.ResponseWriter, r *http.Request)
	DeleteArticleById(w http.ResponseWriter, r *http.Request)
}

func NewArticleController(svc service.ArticleService) ArticleController {
	return &articleController{
		articleService: svc,
	}
}

func (c *articleController) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var post model.Article

	helper.ReadFromRequestBody(r, &post)

	result := c.articleService.CreateArticle(&post)

	response := model.ApiResponse{
		Code: http.StatusCreated,
		Message: "OK",
		Status: "success",
		Data: result,
	}

	helper.WriteToResponseBody(w, response)
}

func (c *articleController) GetArticles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	limit := params["limit"]
	offset := params["offset"]

	defaultLimit, errLimit := strconv.Atoi(limit)
	helper.PanicIfError(errLimit)
	defaultOffset, errOfsset := strconv.Atoi(offset)
	helper.PanicIfError(errOfsset)

	results := c.articleService.GetArticles(defaultLimit,defaultOffset)

	response := model.ApiResponse{
		Code: http.StatusOK,
		Status: "success",
		Message: "List All Posts",
		Data: results,
	}

	helper.WriteToResponseBody(w, response)
}

func (c *articleController) GetArticleById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paramId := params["id"]

	postId, errConv := strconv.Atoi(paramId)

	helper.PanicIfError(errConv)

	post := c.articleService.GetArticleById(postId)

	response := model.ApiResponse{
		Code: http.StatusOK,
		Status: "success",
		Message: "Get Post By Id",
		Data: post,
	}

	helper.WriteToResponseBody(w, response)
	
}

func (c *articleController) UpdateArticleById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paramId := params["id"]

	var post model.Article

	helper.ReadFromRequestBody(r, &post)

	postId, errConv := strconv.Atoi(paramId)

	helper.PanicIfError(errConv)

	postUpdated := c.articleService.UpdateArticleById(postId, &post)

	response := model.ApiResponse{
		Code: http.StatusOK,
		Status: "success",
		Message: "OK",
		Data: postUpdated,
	}

	helper.WriteToResponseBody(w, response)
}

func (c *articleController) DeleteArticleById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paramId := params["id"]
	postId, errConv := strconv.Atoi(paramId)

	helper.PanicIfError(errConv)

	errDelete := c.articleService.DeleteArticleById(postId)

	helper.PanicIfError(errDelete)

	response := model.ApiResponse{
		Code: http.StatusOK,
		Status: "success",
		Message: "Post save to thrash",
	}

	helper.WriteToResponseBody(w, response)
}