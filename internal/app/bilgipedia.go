package app

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

type Route interface {
	Method() string
	Path() string
	HandlerFunc() func(echo.Context) error
}

func InitializeTheServer(port string, routes []Route, authMW echo.MiddlewareFunc) {

	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.Recover())
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

	if err := e.StartTLS(":443", "server.crt", "server.key"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
