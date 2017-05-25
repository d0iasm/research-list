package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"../models"
	// "go-echo-vue/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetPapers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPapers(db))
	}
}

func PutPaper(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var paper models.Paper

		c.Bind(&paper)

		id, err := models.PutPaper(db, paper.Name)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

func DeletePaper(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := models.DeletePaper(db, id)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
