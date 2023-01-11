package app

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme"
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

	autoTLSManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
		Cache: autocert.DirCache("/var/www/.cache"),
		//HostPolicy: autocert.HostWhitelist("<DOMAIN>"),
	}
	s := http.Server{
		Addr:    ":443",
		Handler: e, // set Echo as handler
		TLSConfig: &tls.Config{
			//Certificates: nil, // <-- s.ListenAndServeTLS will populate this field
			GetCertificate: autoTLSManager.GetCertificate,
			NextProtos:     []string{acme.ALPNProto},
		},
		//ReadTimeout: 30 * time.Second, // use custom timeouts
	}

	if err := s.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}

}
