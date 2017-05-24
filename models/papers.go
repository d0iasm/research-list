package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Paper struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PaperCollection struct {
	Papers []Paper `json:"items"`
}

func GetPapers(db *sql.DB) PaperCollection {
	sql := "SELECT * FROM papers"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := PaperCollection{}
	for row.Next() {
		paper := Paper{}
		err2 := rows.Scan(&paper.Id, &paper.Name)

		if err2 != nil {
			panic(err2)
		}

		result.Papers = append(result.Papers, paper)
	}
	return result
}
