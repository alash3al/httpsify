package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(getAvailableHosts()...)
	e.AutoTLSManager.Cache = autocert.DirCache(*flagAutocertCacheDir)

	e.Use(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if *flagSendXSecuredBy {
				c.Response().Header().Set("X-Secured-By", "https://github.com/alash3al/httpsify")
			}
			return next(c)
		}
	})

	e.Any("/*", handler)

	errChan := make(chan error)

	go (func() {
		errChan <- e.Start(*flagHTTPAddr)
	})()

	go (func() {
		errChan <- e.StartAutoTLS(*flagHTTPSAddr)
	})()

	log.Fatal(map[string]interface{}{
		"message": <-errChan,
	})
}
