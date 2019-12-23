package main

import (
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func parseHostsFile(filename string) (map[string]*echo.Echo, error) {
	fd, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	finfo, err := fd.Stat()
	if err != nil {
		return nil, err
	}

	if finfo.Size() == 0 {
		return map[string]*echo.Echo{}, nil
	}

	var rawdata map[string][]string

	if err := json.NewDecoder(fd).Decode(&rawdata); err != nil {
		return nil, err
	}

	compiled := map[string]*echo.Echo{}

	for host, upstreams := range rawdata {
		if len(upstreams) < 1 {
			return nil, errors.New("no upstreams for: " + host)
		}

		targets := []*middleware.ProxyTarget{}
		for _, upstream := range upstreams {
			if !strings.HasPrefix(upstream, "http://") || !strings.HasPrefix(upstream, "https://") {
				upstream = "http://" + upstream
			}

			u, err := url.Parse(upstream)
			if err != nil {
				return nil, err
			}

			targets = append(targets, &middleware.ProxyTarget{
				URL: u,
			})
		}

		e := echo.New()
		e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

		compiled[host] = e
	}

	return compiled, nil
}

func watchHostsChanges(filename string, fn func()) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	watcher.Add(filename)

	for {
		select {
		case <-watcher.Events:
			fn()
		}
	}
}

func getAvailableHosts() []string {
	result := []string{}

	for k := range hosts.Load().(map[string]*echo.Echo) {
		result = append(result, k)
	}

	return result
}
