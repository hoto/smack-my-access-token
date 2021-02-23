package main

import (
	"fmt"
	"github.com/hoto/smack-my-access-token/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.ParseArgsAndFlags()
	secrets := loadSecrets()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/docusign", generateDocusignJWT(secrets))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
