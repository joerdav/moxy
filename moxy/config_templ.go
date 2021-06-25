// Code generated by templ DO NOT EDIT.

package moxy

import "github.com/a-h/templ"
import "context"
import "io"
import "github.com/joe-davidson1802/moxy/config"

func ConfigTemplate(c *config.Config, message string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		_, err = io.WriteString(w, "<turbo-frame")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " id=\"table\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		var templCSSClassess templ.CSSClasses = templ.Classes(transparent(), green(), templ.Class("fade-out"))
		err = templ.RenderCSS(ctx, w, templCSSClassess)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<h6")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(templCSSClassess.String()))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(message))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</h6>")
		if err != nil {
			return err
		}
		err = UpstreamTableTemplate(c).Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</turbo-frame>")
		if err != nil {
			return err
		}
		return err
	})
}

