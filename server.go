package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

}

func main() {
	db := initDB("storage.db")
	migrate(db)

	e := echo.New()
	e.GET("/", getRoot)
	e.Logger.Fatal(e.Start(":1323"))
}
