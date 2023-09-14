package main

import (
	"net/http"

	"github.com/shoelfikar/golang_backend_api/app"
	"github.com/shoelfikar/golang_backend_api/helper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	app.NewViperLoad()

	DB := app.NewDB()

	router := app.NewRouter(DB)


	driver, err := migrate.New("file://db/migrations", helper.GetViperEnvVariable("MIGRATE_URL"))
    if err != nil {
		helper.DefaultLoggingError(err.Error())
    }

    err = driver.Up()
    if err != nil && err != migrate.ErrNoChange {
        helper.DefaultLoggingError(err.Error())
    }

	server := http.Server{
		Addr: ":" + helper.GetViperEnvVariable("PORT"),
		Handler: router,
	}

	server.ListenAndServe()
}