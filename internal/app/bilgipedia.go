package app

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Route interface {
	Method() string
	Path() string
	HandlerFunc() func(echo.Context) error
}

func InitializeTheServer(port string, routes []Route, authMW echo.MiddlewareFunc) {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowCredentials, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowMethods, echo.HeaderAccessControlAllowMethods},
	}))

	// INJECTING THE ROUTES
	for _, route := range routes {
		if route.Method() == echo.POST && !strings.HasPrefix(route.Path(), "/sign") {
			e.Add(route.Method(), route.Path(), route.HandlerFunc(), authMW)
		} else {
			e.Add(route.Method(), route.Path(), route.HandlerFunc())
		}

	}

	e.Logger.Fatal(e.Start(":8081"))
}
