HTTPSify 
========
A `Let'sEncrypt` based reverse proxy, that will automatically generate &amp; renew valid `ssl` certs for your domains, it also enables the `http/2` protocol by default, and uses `roundrobin` as an algorithm to loadbalance the incoming requests between multiple `upstreams`, as well as redirecting the traffic from `http` traffic to `https` just if you enabled the flag `--redirect`.

NOTES
=======
> HTTPSify only supports `http-01` challenge because [Let's Encrypt disables TLS-SNI-01 validation](http://www.zdnet.com/article/lets-encrypt-disables-tls-sni-01-validation/)

# Quick Start

### # Using Docker
> Just run the following and then have fun !!
```bash
$ docker run --network host -v ~/.httpsify:/.httpsify -p 443:443 alash3al/httpsify
```

## # From Binaries
> Go to [releases page](https://github.com/alash3al/httpsify/releases)

### # Building from source
> You must have the `Go` environment installed
```bash
$ go get -u github.com/alash3al/httpsify
```

### # Configurations
> Goto your `$HOME` Directory and edit the `hosts.json` to something like this
```json
{
	"example1.com": ["http://localhost:9080"],
	"example2.com": ["http://localhost:8080", "http://localhost:8081"]
}
```
> As you see, the configuration file accepts a `JSON` object/hashmap of `domain` -> `upstreams`,
and yes, it can loadbalance the requests between multiple upstreams using `roundrobin` algorithm.

> Also You don't need to restart the server to reload the configurations, because `httpsify` automatically watches the
configurations file and reload it on any change.

# License
> The MIT License (MIT)

> Copyright (c) 2016 Mohammed Al Ashaal

> Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
