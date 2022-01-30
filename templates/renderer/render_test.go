package renderer_test

import (
	"bytes"
	"jitsusama/lgwt/templates/renderer"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
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
		approvals.VerifyString(t, buf.String())
	})
}
