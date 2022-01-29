package posts

import (
	"bufio"
	"fmt"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
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

	scanner := bufio.NewScanner(file)

	getMeta := func(name string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), fmt.Sprintf("%s: ", name))
	}

	return Post{
		Title:       getMeta("Title"),
		Description: getMeta("Description"),
	}, nil
}
