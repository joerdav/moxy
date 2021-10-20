package moxy

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-davidson1802/moxy/config"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

//go:embed css
var content embed.FS

func NewMoxy(c *config.Config) http.Handler {
	r := mux.NewRouter()

	fs := logRequest(http.FileServer(http.FS(content)))

	r.
		Handle(c.MoxyPath+"/config", ConfigHandler{config: c}).Methods("GET")
	r.
		Handle(c.MoxyPath+"/config", AddConfigHandler{config: c}).Methods("POST")
	r.
		Handle(c.MoxyPath+"/delete", DeleteConfigHandler{config: c}).Methods("POST")
	r.
		Handle(c.MoxyPath+"/move", MoveConfigHandler{config: c}).Methods("POST")
	r.
		Handle(c.MoxyPath+"/host", UpdateHostHandler{config: c}).Methods("PUT")
	r.
		Handle(c.MoxyPath+"/pathprefix", UpdatePathPrefixHandler{config: c}).Methods("PUT")
	r.
		PathPrefix(c.MoxyPath + "/css").
		Handler(logRequest(http.StripPrefix(c.MoxyPath, fs)))
	r.
		Handle(c.MoxyPath, HomeHandler{config: c})

	return r
}
