package renderer

import (
	"embed"
	"html/template"
	"io"
)

//go:embed "*.tmpl"
var templates embed.FS

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	template, err := template.ParseFS(templates, "*.html.tmpl")
	if err != nil {
		return err
	}
	if err := template.Execute(w, p); err != nil {
		return err
	}
	return nil
}
