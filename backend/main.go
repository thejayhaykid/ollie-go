package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!"})
	})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("Middleware 1")
			fmt.Println("Request URI: ", c.Request().RequestURI)
			return next(c)
		}
	})

	e.Logger.Fatal(e.Start(":8080"))
}
