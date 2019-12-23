package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func preloadCerts() {
	hosts := hosts.Load().(map[string]*echo.Echo)

	for k := range hosts {
		go (func(h string) {
			resp, err := http.Get("http://" + h)
			if err != nil {
				log.Println(err.Error())
				return
			}
			defer resp.Body.Close()
			log.Printf("=> preloading certs result for %s -> %s\n", h, resp.Status)
		})(k)
	}
}
