package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", getRoot)
	e.Logger.Fatal(e.Start(":1323"))
}
