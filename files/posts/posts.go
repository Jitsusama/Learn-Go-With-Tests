package posts

import (
	"io/fs"
)

type Post struct{}

func NewPostsFromFs(filesystem fs.FS) []Post {
	dir, _ := fs.ReadDir(filesystem, ".")

	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts
}
