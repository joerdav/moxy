package moxy

import (
	"net/http"

	"github.com/joe-davidson1802/moxy/config"
)

type HomeHandler struct {
	config *config.Config
}

func (h HomeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/html")

	HomeTemplate(h.config).Render(req.Context(), res)
}
