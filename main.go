package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	_ "pkhood-backend/docs"
	"pkhood-backend/src/cmd"
	"pkhood-backend/src/settings"
)

// @title Space Swagger
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name Furkan AYDIN
// @contact.email furkan139aydin@gmail.com

// @host localhost:5000
// @BasePath /api
// @securityDefinitions.apikey PToken
// @in header
// @name Authorization
func main() {

	settings.InitEnvironment()

	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/api")

	g.GET("/swagger/*", echoSwagger.WrapHandler)

	cmd.ExecuteUser(g)

	e.Logger.Fatal(e.Start(":5000"))
}
