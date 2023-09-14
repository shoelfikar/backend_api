package app

import (
	"database/sql"
	"time"

	"github.com/shoelfikar/golang_backend_api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", helper.GetViperEnvVariable("DB_URL"))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
