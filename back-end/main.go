package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", returnValue)

	log.Println("Server is available at http://localhost:8000")
	e.Start(":8000")
}

func returnValue(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
