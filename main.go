package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, fi os.FileInfo, err error) error {
	if !fi.IsDir() {
		return filepath.SkipDir
	}
	switch filepath.Base(fi.Name()) {
	case ".":
		return filepath.SkipDir
	case ".git":
		return errors.New("end")
	}

	if filepath.Base(fi.Name()) == ".git" {
		return errors.New("end")
	}
	fmt.Println(path)
	return nil
}

func main() {
	root := "."
	for {
		err := filepath.Walk(root, walkFunc)
		if err != nil {
			os.Exit(1)
		}
		root = filepath.Join(root, "..")
	}
}
