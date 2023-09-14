package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/shoelfikar/golang_backend_api/helper"
	"github.com/shoelfikar/golang_backend_api/model"
	"github.com/shoelfikar/golang_backend_api/repository"
)

type articleService struct {
	articleRepo repository.ArticleRepository
	validate    *validator.Validate
}

type ArticleService interface {
	CreateArticle(post *model.Article) *model.Article
	GetArticles(limit int, offset int) []*model.Article
	GetArticleById(postId int) *model.Article
	UpdateArticleById(postId int, post *model.Article) *model.Article
	DeleteArticleById(postId int) error
}

func NewArticleService(repo repository.ArticleRepository, validate *validator.Validate) ArticleService {
	return &articleService{
		articleRepo: repo,
		validate: validate,
	}
}

func (a *articleService) CreateArticle(post *model.Article) *model.Article {
	err := a.validate.Struct(post)

	helper.PanicIfError(err)

	return a.articleRepo.CreateArticle(post)
}

func (a *articleService) GetArticles(limit int, offset int) []*model.Article {
	defaultLimit := 10
	defaultOffset := 0

	if limit != 0 && offset != 0 {
		defaultLimit = limit
		defaultOffset = offset
	}

	return a.articleRepo.GetArticles(defaultLimit, defaultOffset)
}

func (a *articleService) GetArticleById(postId int) *model.Article {
	return a.articleRepo.GetArticleById(postId)
}

func (a *articleService) UpdateArticleById(postId int, post *model.Article) *model.Article {
	err := a.validate.Struct(post)

	helper.PanicIfError(err)

	return a.articleRepo.UpdateArticleById(postId, post)
}

func (a *articleService) DeleteArticleById(postId int) error {
	return a.articleRepo.DeleteArticleById(postId)
}
