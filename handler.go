package main

import "github.com/labstack/echo/v4"

func handler(c echo.Context) error {
	req := c.Request()
	res := c.Response()

	hosts := hosts.Load().(map[string]*echo.Echo)
	host := hosts[req.Host]

	if host == nil {
		return echo.ErrNotFound
	}

	host.ServeHTTP(res, req)

	return nil
}
