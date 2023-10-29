package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})
	e.Logger.Fatal(e.Start(":" + port))
}
