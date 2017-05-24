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

func PutPaper(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO papers(name) VALUES(?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(name)

	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

func DeletePaper(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM papers WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)

	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}
