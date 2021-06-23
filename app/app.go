package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/joe-davidson1802/moxy/config"
	"github.com/joe-davidson1802/moxy/proxy"
)

func Run() {
	var c config.Config

	file := "defaults.toml"

	fmt.Println("Loading Defaults...")

	if _, err := toml.DecodeFile(file, &c); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.
		PathPrefix("/").
		Handler(proxy.New(c))

	fmt.Println("Listening: :80")

	err := http.ListenAndServe(":80", r)

	log.Fatal(err)
}
