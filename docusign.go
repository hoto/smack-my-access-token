package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Message string `json:"name"`
}

func generateDocusignJWT(c echo.Context) error {
	r := &Response{
		Message: "Hello world!",
	}
	return c.JSON(http.StatusOK, r)
}
