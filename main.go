package main

import (
	"fmt"
	"github.com/hoto/smack-my-access-token/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.ParseArgsAndFlags()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", generateDocusignJWT)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
