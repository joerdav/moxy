// Code generated by templ DO NOT EDIT.

package moxy

import "github.com/a-h/templ"
import "context"
import "io"
import "github.com/joe-davidson1802/moxy/config"

func UpstreamUrlTemplate(c *config.Config, id string, url string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<form")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " action=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(c.MoxyPath + "/url"))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " method=\"PUT\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<input")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"id\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " value=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(id))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " size=\"50\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " hidden=\"true\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "/>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<input")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"url\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " value=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(url))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " size=\"50\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "/>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<button")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " type=\"submit\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString("Update"))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</button>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</form>")
		if err != nil {
			return err
		}
		return err
	})
}

func UpstreamPathPrefixTemplate(c *config.Config, id string, pathprefix string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<form")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " action=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(c.MoxyPath + "/pathprefix"))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " method=\"PUT\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<input")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"id\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " value=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(id))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " size=\"50\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " hidden=\"true\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "/>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<input")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"path\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " value=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(pathprefix))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " size=\"50\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "/>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<button")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " type=\"submit\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString("Update"))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</button>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</form>")
		if err != nil {
			return err
		}
		return err
	})
}
