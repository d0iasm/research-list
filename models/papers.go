package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Paper struct {
	ID int `json:"id"`
	Title string `json:"name"`
	Author string `json:"name"`
	Journal string `json:"name"`
	Year int16 `json:"name"`
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
	for rows.Next() {
		paper := Paper{}
		//err2 := rows.Scan(&paper.ID, &paper.Name)
		err2 := rows.Scan(&paper.ID, &paper.Title)

		if err2 != nil {
			panic(err2)
		}

		result.Papers = append(result.Papers, paper)
	}
	return result
}

func PutPaper(db *sql.DB, title string) (int64, error) {
	//sql := "INSERT INTO papers(name) VALUES(?)"
	sql := "INSERT INTO papers(title) VALUES(?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	//result, err2 := stmt.Exec(name)
	result, err2 := stmt.Exec(title)

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
