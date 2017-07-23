package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Paper struct {
	ID int `json:"id"`
	Title string `json:"name"`
	Author string `json:"author"`
	Journal string `json:"journal"`
	Year int16 `json:"year"`
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
		err2 := rows.Scan(&paper.ID, &paper.Title, &paper.Author, &paper.Journal, &paper.Year)

		if err2 != nil {
			panic(err2)
		}

		result.Papers = append(result.Papers, paper)
	}
	return result
}

// func PutPaper(db *sql.DB, title string) (int64, error) {
func PutPaper(db *sql.DB, title, author, journal string, year int16) (int64, error) {
	//sql := "INSERT INTO papers(name) VALUES(?)"
	sql := "INSERT INTO papers(title, author, journal, year) VALUES(?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(sql)
	// result, err2 := stmt.Exec(title, author, journal, year)

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
