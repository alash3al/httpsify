package main

import (
	"flag"
	"path"
	"sync/atomic"

	"github.com/mitchellh/go-homedir"
)

var (
	homeDir, _ = homedir.Dir()
	hosts      = atomic.Value{}
)

var (
	flagHTTPSAddr        = flag.String("https", ":443", "the port to listen for https requests on, it is recommended to leave it as is")
	flagHTTPAddr         = flag.String("http", ":80", "the port to listen for http requests on, it is recommended to leave it as is")
	flagAutocertCacheDir = flag.String("certs", path.Join(homeDir, ".httpsify/certs"), "the certs directory")
	flagHostsFile        = flag.String("hosts", path.Join(homeDir, ".httpsify/hosts.json"), "the file containing hosts mappings to upstreams")
	flagSendXSecuredBy   = flag.Bool("x-secured-by", true, "whether to enable x-secured-by header or not")
)
