package renderer_test

import (
	"bytes"
	"jitsusama/lgwt/templates/renderer"
	"testing"
)

func TestRender(t *testing.T) {
	examplePost := renderer.Post{
		Title:       "Hello World",
		Description: "This is an introduction to my blog!",
		Tags:        []string{"welcome", "wall-of-text"},
		Body:        "This is a post.",
	}
	t.Run("converts post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := renderer.Render(&buf, examplePost); err != nil {
			t.Fatal(err)
		}
		actual := buf.String()

		expected := `<h1>Hello World</h1>
<p>This is an introduction to my blog!</p>
Tags: <ul><li>welcome</li><li>wall-of-text</li></ul>`
		if actual != expected {
			t.Errorf("got %q want %q", actual, expected)
		}
	})
}
