package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "embed"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/joe-davidson1802/moxy/config"
	"github.com/joe-davidson1802/moxy/moxy"
	"github.com/joe-davidson1802/moxy/proxy"
)

//go:embed defaults.toml
var cfile string

func Run() {
	var c config.Config
	var cstring = cfile

	if len(os.Args) >= 2 {
		fmt.Printf("Loading config from %s\n", os.Args[1])
		b, err := os.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		cstring = string(b)
	}
	if _, err := toml.Decode(cstring, &c); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.
		PathPrefix(c.MoxyPath).
		Handler(moxy.NewMoxy(&c))

	r.
		PathPrefix("/").
		Handler(proxy.New(&c))

	fmt.Println("Listening: :80")

	err := http.ListenAndServe(":80", r)

	log.Fatal(err)
}
