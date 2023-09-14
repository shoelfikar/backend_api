package repository

import "github.com/shoelfikar/golang_backend_api/model"

type ArticleRepository interface {
	CreateArticle(post *model.Article) *model.Article
	GetArticles(limit int, offset int) []*model.Article
	GetArticleById(postId int) *model.Article
	UpdateArticleById(postId int, post *model.Article) *model.Article
	DeleteArticleById(postId int) error
}