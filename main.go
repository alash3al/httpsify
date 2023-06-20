package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(getAvailableHosts()...)
	e.AutoTLSManager.Cache = autocert.DirCache(*flagAutocertCacheDir)
	e.AutoTLSManager.Client = &acme.Client{
		DirectoryURL: *flagACMEDirectory,
		UserAgent:    "https://github.com/alash3al/httpsify",
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if *flagSendXSecuredBy {
				c.Response().Header().Set("X-Secured-By", "https://github.com/alash3al/httpsify")
			}

			hosts := hosts.Load().(map[string]*echo.Echo)
			host := hosts[c.Request().Host]

			if host == nil {
				return echo.ErrNotFound
			}

			return echo.WrapHandler(host)(c)
		}
	})

	errChan := make(chan error)

	go (func() {
		errChan <- http.ListenAndServe(*flagHTTPAddr, e.AutoTLSManager.HTTPHandler(nil))
	})()

	go (func() {
		errChan <- e.StartAutoTLS(*flagHTTPSAddr)
	})()

	log.Fatal(map[string]interface{}{
		"message": <-errChan,
	})
}
