package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Message string `json:"name"`
}

func generateDocusignJWT(secrets Secrets) func(echo.Context) error {
	return func(context echo.Context) error {
		response := &Response{
			Message: fmt.Sprintf("Smack my username=%s", secrets.Docusign.IntegrationKey),
		}
		return context.JSON(http.StatusOK, response)
	}
}

