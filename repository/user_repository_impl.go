package repository

import (
	"context"
	"database/sql"

	"github.com/shoelfikar/golang_backend_api/helper"
	"github.com/shoelfikar/golang_backend_api/model"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

func (u *userRepositoryImpl) Create(usr *model.User) *model.User {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	ctx := context.Background()
	
	
	query := "INSERT INTO users (username,password,email,status,created_by) VALUES(?,?,?,?,?)"
	
	res, err := tx.ExecContext(ctx, query, &usr.Username, &usr.Password, &usr.Email, &usr.Status, "system")
	
	if err != nil {
		tx.Rollback()
		helper.PanicIfError(err)
	}
	
	lastId, _ := res.LastInsertId()
	
	usr.Id = int(lastId)
	
	tx.Commit()
	
	return usr
}

func (u *userRepositoryImpl) GetUserById(userId int) *model.User {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	ctx := context.Background()

	var user model.User

	query := "SELECT id,username,email,password,status,created_at,created_by WHERE id = ?"

	rowErr := tx.QueryRowContext(ctx, query, &userId).Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.Status,&user.CreatedAt,&user.CreatedBy)

	if rowErr != nil {
		if rowErr == sql.ErrNoRows {
			panicErr := model.NotFoundError{Error: rowErr.Error()}
			panic(panicErr)
		}
		helper.PanicIfError(rowErr)
	}

	return &user
}

func (u *userRepositoryImpl) GetUserByEmail(login *model.LoginRequest) *model.User {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	ctx := context.Background()

	var user model.User

	query := "SELECT id,username,email,password,status,created_at,created_by FROM users WHERE email = ? "


	rows := tx.QueryRowContext(ctx, query, &login.Email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.CreatedBy)

	if rows != nil {
		if rows == sql.ErrNoRows {
			panicErr := model.NotFoundError{Error: rows.Error()}
			panic(panicErr)
		}
		
		helper.PanicIfError(rows)
	}

	return &user

	
}