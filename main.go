package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	current = "."
	parent  = ".."
)

type skip struct {
	message string // skip message
	code    int    // staus code
	cause   error  // the original error object
	end     bool   // the end of scanning
}

func (r skip) Error() string {
	return r.message
}

func walk(path string, fi os.FileInfo, err error) error {
	if !fi.IsDir() {
		return filepath.SkipDir
	}

	// About skipping in a special case
	switch filepath.Base(fi.Name()) {
	case ".":
		return skip{
			message: "dot dir (.) should be skipped",
			code:    0,
			cause:   filepath.SkipDir,
			end:     false,
		}
	case ".git":
		return skip{
			message: ".git directory was found",
			code:    0,
			cause:   nil,
			end:     true, // stop scanning
		}
	}

	// TODO: scan through the directory where the .git directory is located (means the top directory)

	fmt.Println(filepath.Clean(path))
	return nil
}

func main() {
	cwd := current
	for {
		err := filepath.Walk(cwd, walk)
		if err != nil {
			switch err := err.(type) {
			case skip:
				if err.end {
					os.Exit(err.code)
				}
			default:
				panic(err)
			}
		}
		cwd = filepath.Join(cwd, parent)
	}
}
