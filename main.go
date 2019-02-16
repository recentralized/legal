package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

const (
	srcDir = "src"
	dstDir = "dist"
)

func main() {

	paths, err := filepath.Glob(srcDir + "/*.md")
	if err != nil {
		fmt.Printf("Failed to list files: %v\n", err)
		os.Exit(1)
	}

	for _, path := range paths {
		out := pathFor(path)
		fmt.Printf("Rendering %s to %s\n", path, out)
		store(out, render(path))
	}
}

func pathFor(p string) string {
	base := path.Base(p)
	ext := path.Ext(p)
	base = strings.Replace(base, ext, ".html", 1)
	return path.Join(dstDir, base)
}

func render(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	input, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		os.Exit(1)
	}

	return blackfriday.Run(input)
}

func store(path string, data []byte) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		os.Exit(1)
	}

	if _, err := f.Write(data); err != nil {
		fmt.Printf("Failed to write file: %v\n", err)
		os.Exit(1)
	}
}
