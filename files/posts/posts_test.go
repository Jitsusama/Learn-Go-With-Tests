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
Description: Hello World!
Tags: hello
---
Hello to the world!
Peace`)},
		"02-yesterday.md": {Data: []byte(`Title: Yesterday
Description: Was quite the day.
Tags: diary, boring
---
It was a day, a boring day.
B
L
M`)},
	}

	got, err := posts.NewPostsFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(got) != len(fs) {
		t.Errorf("got %d posts want %d posts", len(got), len(fs))
	}

	expected := posts.Post{
		Title:       "Yesterday",
		Description: "Was quite the day.",
		Tags:        []string{"diary", "boring"},
		Body: `It was a day, a boring day.
B
L
M`,
	}
	if !postsEqual(got[1], expected) {
		t.Errorf("got %+v want %+v", got[1], expected)
	}
}

func postsEqual(actual posts.Post, expected posts.Post) bool {
	return reflect.DeepEqual(actual, expected)
}
