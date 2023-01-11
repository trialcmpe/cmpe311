package app

import (
	"crypto/tls"
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

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	s := http.Server{
		Addr:      ":443",
		Handler:   e, // set Echo as handler
		TLSConfig: cfg,
		//ReadTimeout: 30 * time.Second, // use custom timeouts
	}

	if err := s.ListenAndServeTLS("server.crt", "server.key"); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}

}
