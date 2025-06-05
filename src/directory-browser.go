package main

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	path := filepath.Dir(exe)
	result := searchDirectory(path, make([]string, 0), "")
	os.WriteFile("vyhledatkovano.csv", []byte(result), fs.ModeDir)
}

func searchDirectory(
	directory string,
	fullPath []string,
	contents string,
) string {
	entries, err := os.ReadDir(directory)
	var result = contents

	if err != nil {
		panic(err)
	}

	for _, e := range entries {
		name := e.Name()
		nextFullPath := append(fullPath, name)
		nextName := path.Join(nextFullPath...)
		if e.IsDir() {
			result = searchDirectory(nextName, nextFullPath, result)
		} else {
			split := strings.ReplaceAll(nextName, "/", ";")
			result = result + split + "\n"
		}
	}

	return result
}
