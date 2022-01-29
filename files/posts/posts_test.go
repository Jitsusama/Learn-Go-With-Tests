package posts_test

import (
	"jitsusama/lgwt/files/posts"
	"testing"
	"testing/fstest"
)

func TestNewPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts, err := posts.NewPostsFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts want %d posts", len(posts), len(fs))
	}
}
