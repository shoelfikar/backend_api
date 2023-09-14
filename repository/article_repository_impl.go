package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/shoelfikar/golang_backend_api/helper"
	"github.com/shoelfikar/golang_backend_api/model"
)

type articleRepository struct {
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{
		DB: db,
	}
}

func (a *articleRepository) CreateArticle(post *model.Article) *model.Article {
	tx, err := a.DB.Begin()
	ctx := context.Background()

	helper.PanicIfError(err)

	query := "INSERT INTO posts (title,content,category,status) VALUES(?,?,?,?)"


	result, errQuery := tx.ExecContext(ctx, query, &post.Title, &post.Content, &post.Category, &post.Status)

	if errQuery != nil {
		tx.Rollback()
		helper.PanicIfError(errQuery)
	}

	lastId, _ := result.LastInsertId()

	post.Id = int(lastId)

	tx.Commit()

	return post
}

func (a *articleRepository) GetArticles(limit int, offset int) []*model.Article {
	DB := a.DB

	ctx := context.Background()

	var posts []*model.Article

	query := "SELECT title,content,category,status FROM posts ORDER BY id DESC LIMIT ? OFFSET ?"



	rows , errQuery := DB.QueryContext(ctx, query, &limit, &offset)

	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			panicErr := model.NotFoundError{Error: errQuery.Error()}
			panic(panicErr)
		}

		helper.PanicIfError(errQuery)
	}

	for rows.Next() {
		post := model.Article{}

		rows.Scan(&post.Title, &post.Content, &post.Category, &post.Status)

		posts = append(posts, &post)

	}

	return posts
}

func (a *articleRepository) GetArticleById(postId int) *model.Article {
	DB := a.DB

	ctx := context.Background()

	var post model.Article

	query := "SELECT title,content,category,status FROM posts WHERE id = ? "

	errQuery := DB.QueryRowContext(ctx, query, &postId).Scan(&post.Title, &post.Content, &post.Category, &post.Status)

	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			panicErr := model.NotFoundError{Error: errQuery.Error()}
			panic(panicErr)
		}
		
		helper.PanicIfError(errQuery)
	}

	return &post
}

func (a *articleRepository) UpdateArticleById(postId int, post *model.Article) *model.Article {
	tx, err := a.DB.Begin()

	helper.PanicIfError(err)

	ctx := context.Background()


	query := "UPDATE posts SET title = ?, content = ?, category = ?, status = ? WHERE id = ? "

	result, errQuery := tx.ExecContext(ctx, query, &post.Title, &post.Content, &post.Category, &post.Status, &postId)

	if errQuery != nil {
		tx.Rollback()
		if errQuery == sql.ErrNoRows {
			panicErr := model.NotFoundError{Error: errQuery.Error()}
			panic(panicErr)
		}

		helper.PanicIfError(errQuery)
	}

	rowAffected, _ := result.RowsAffected()

	if rowAffected <= 0 {
		tx.Rollback()
		noAffected := errors.New("no data updated")
		helper.PanicIfError(noAffected)
	}

	tx.Commit()

	return post
}

func (a *articleRepository) DeleteArticleById(postId int) error {
	tx, err := a.DB.Begin()

	ctx := context.Background()

	if err != nil {
		return err
	}

	query := "UPDATE posts SET status = 'Thrash' WHERE id = ? "

	result , errQuery := tx.ExecContext(ctx, query, &postId)

	if errQuery != nil {
		tx.Rollback()
		return errQuery
	}

	rowAffected, _ := result.RowsAffected()

	if rowAffected <= 0 {
		tx.Rollback()
		noAffected := errors.New("no data updated")
		return noAffected
	}

	tx.Commit()

	return nil
}