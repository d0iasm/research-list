package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetPapers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "papers")
	}
}

func PutPaper(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{
			"created": 123,
		})
	}
}

func DeletePaper(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
