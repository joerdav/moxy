package moxy

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-davidson1802/moxy/config"
)

func NewMoxy(c *config.Config) http.Handler {
	r := mux.NewRouter()

	r.
		Handle("/moxy/config", ConfigHandler{config: c}).Methods("GET")
	r.
		Handle("/moxy/config", AddConfigHandler{config: c}).Methods("POST")
	r.
		Handle("/moxy/delete", DeleteConfigHandler{config: c}).Methods("POST")
	r.
		Handle("/moxy/move", MoveConfigHandler{config: c}).Methods("POST")
	r.
		Handle("/moxy/url", UpdateUrlHandler{config: c}).Methods("PUT")
	r.
		Handle("/moxy/pathprefix", UpdatePathPrefixHandler{config: c}).Methods("PUT")
	r.
		Handle("/moxy", HomeHandler{config: c})

	return r
}
