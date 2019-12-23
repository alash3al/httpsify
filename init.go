package main

import (
	"flag"
	"log"
	"os"
)

func init() {
	flag.Parse()

	os.MkdirAll(*flagAutocertCacheDir, 0755)

	{
		resultedHosts, err := parseHostsFile(*flagHostsFile)
		if err != nil {
			log.Fatal(err.Error())
		}

		hosts.Store(resultedHosts)
	}

	{
		go watchHostsChanges(*flagHostsFile, func() {
			log.Println("⇨ reloading hosts file ...")
			resultedHosts, err := parseHostsFile(*flagHostsFile)
			if err != nil {
				log.Println("⇨ ", err.Error())
				return
			}

			hosts.Store(resultedHosts)
		})
	}

	{
		go preloadCerts()
	}
}
