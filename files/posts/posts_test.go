package posts_test

import (
	"jitsusama/lgwt/files/posts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}

	got, err := posts.NewPostsFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(got) != len(fs) {
		t.Errorf("got %d posts want %d posts", len(got), len(fs))
	}

	expected := posts.Post{Title: "Post 1"}
	if !reflect.DeepEqual(got[0], expected) {
		t.Errorf("got %+v want %+v", got[0], expected)
	}
}
