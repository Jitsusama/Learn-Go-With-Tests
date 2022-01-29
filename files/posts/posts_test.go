package posts_test

import (
	"jitsusama/lgwt/files/posts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewPosts(t *testing.T) {
	fs := fstest.MapFS{
		"01-hello.md": {Data: []byte(`Title: Hello
Description: Hello World!`)},
		"02-yesterday.md": {Data: []byte(`Title: Yesterday
Description: Was quite the day.`)},
	}

	got, err := posts.NewPostsFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(got) != len(fs) {
		t.Errorf("got %d posts want %d posts", len(got), len(fs))
	}

	expected := posts.Post{Title: "Hello", Description: "Hello World!"}
	if !postsEqual(got[0], expected) {
		t.Errorf("got %+v want %+v", got[0], expected)
	}
}

func postsEqual(actual posts.Post, expected posts.Post) bool {
	return reflect.DeepEqual(actual, expected)
}
