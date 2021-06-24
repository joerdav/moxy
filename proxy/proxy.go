package proxy

import (
	"fmt"

	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/joe-davidson1802/moxy/config"
)

func New(c *config.Config) ProxyHandler {
	return ProxyHandler{
		config: c,
	}
}

type ProxyHandler struct {
	config *config.Config
}

func (h ProxyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	source := req.URL.String()

	url, _ := url.Parse(h.config.GetDestination(req.RequestURI))

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Host = url.Host

	fmt.Printf("%s -> %s\n", source, req.URL.String())

	proxy.ServeHTTP(res, req)
}
