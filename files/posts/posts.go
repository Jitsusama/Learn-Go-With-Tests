package posts

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFs(filesystem fs.FS) ([]Post, error) {
	filenames, err := getFilenames(filesystem)
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, filename := range filenames {
		post, err := getPost(filesystem, filename)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func getFilenames(filesystem fs.FS) ([]string, error) {
	directory, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, file := range directory {
		filenames = append(filenames, file.Name())
	}
	return filenames, nil
}

func getPost(filesystem fs.FS, filename string) (Post, error) {
	file, err := filesystem.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	d, err := io.ReadAll(file)
	if err != nil {
		return Post{}, err
	}

	return Post{Title: string(d)[7:]}, nil
}
