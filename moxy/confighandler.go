package moxy

import (
	"net/http"
	"strconv"

	"github.com/joe-davidson1802/moxy/config"
	"github.com/joe-davidson1802/moxy/types"
	"github.com/joe-davidson1802/turbo-templ/turbo"
)

type ConfigHandler struct {
	config *config.Config
}

func (h ConfigHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/html")

	content := ConfigTemplate(h.config, " ")

	frame := turbo.TurboFrame(turbo.TurboFrameOptions{Id: "moxy-config", Contents: &content})

	frame.Render(req.Context(), res)
}

type UpdateHostHandler struct {
	config *config.Config
}

func (h UpdateHostHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/html")

	v := req.FormValue("host")
	id := req.FormValue("id")

	if id == "default" {
		h.config.DefaultUpstream = v
	} else {
		i, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		h.config.Upstreams[i].Host = v
	}

	content := ConfigTemplate(h.config, "Updated host")

	frame := turbo.TurboFrame(turbo.TurboFrameOptions{Id: "table", Contents: &content})

	frame.Render(req.Context(), res)
}

type UpdatePathPrefixHandler struct {
	config *config.Config
}

func (h UpdatePathPrefixHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/html")

	id := req.FormValue("id")
	v := req.FormValue("path")

	if id == "default" {
		h.config.DefaultUpstream = v
	} else {
		i, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		h.config.Upstreams[i].PathPrefix = v
	}

	content := ConfigTemplate(h.config, "Updated path")

	frame := turbo.TurboFrame(turbo.TurboFrameOptions{Id: "table", Contents: &content})

	frame.Render(req.Context(), res)
}

type AddConfigHandler struct {
	config *config.Config
}

func (h AddConfigHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/html")

	u := req.FormValue("host")
	v := req.FormValue("pathprefix")

	up := config.Upstream{
		Host:       u,
		PathPrefix: v,
	}

	h.config.Upstreams = append(h.config.Upstreams, up)

	content := ConfigTemplate(h.config, "Added new upstream")

	frame := turbo.TurboFrame(turbo.TurboFrameOptions{Id: "table", Contents: &content})

	frame.Render(req.Context(), res)
}

type MoveConfigHandler struct {
	config *config.Config
}

func (h MoveConfigHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/html")

	id, err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		panic(err)
	}

	dir := types.NewMoveDirection(req.FormValue("direction"))

	if (id == 0 && dir == types.DirectionUp) ||
		(id == len(h.config.Upstreams)-1 && dir == types.DirectionDown) {
		res.WriteHeader(400)
		return
	}

	if dir == types.DirectionUp {
		h.config.Upstreams[id], h.config.Upstreams[id-1] = h.config.Upstreams[id-1], h.config.Upstreams[id]
	} else {
		h.config.Upstreams[id], h.config.Upstreams[id+1] = h.config.Upstreams[id+1], h.config.Upstreams[id]
	}

	content := ConfigTemplate(h.config, "Moved upstream priority")

	frame := turbo.TurboFrame(turbo.TurboFrameOptions{Id: "table", Contents: &content})

	frame.Render(req.Context(), res)
}

type DeleteConfigHandler struct {
	config *config.Config
}

func (h DeleteConfigHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/html")

	id, err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		panic(err)
	}

	h.config.Upstreams = append(h.config.Upstreams[:id], h.config.Upstreams[id+1:]...)

	content := ConfigTemplate(h.config, "Upstream deleted")

	frame := turbo.TurboFrame(turbo.TurboFrameOptions{Id: "table", Contents: &content})

	frame.Render(req.Context(), res)
}
