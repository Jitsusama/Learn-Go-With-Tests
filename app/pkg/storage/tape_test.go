package storage

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestTape(t *testing.T) {
	t.Run("overwrites file contents on write", func(t *testing.T) {
		file, cleanup := createFile(t, "12345")
		defer cleanup()

		tape := &tape{file}
		tape.Write([]byte("abc"))
		file.Seek(0, 0)

		actual, _ := ioutil.ReadAll(file)
		expected := "abc"
		if string(actual) != expected {
			t.Errorf("got %q want %q", actual, expected)
		}
	})
}

func createFile(t testing.TB, contents string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "*.json")
	if err != nil {
		t.Fatalf("temp file creation error: %v", err)
	}
	tmpFile.Write([]byte(contents))

	return tmpFile, func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}
}
