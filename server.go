package main

import (
	"./handlers"
	"database/sql"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/papers", handlers.GetPapers(db))
	e.PUT("/papers", handlers.PutPaper(db))
	e.DELETE("/papers/:id", handlers.DeletePaper(db))

	e.Logger.Fatal(e.Start(":1323"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS papers(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR,
		title VARCHAR NOT NULL
	);
	`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
