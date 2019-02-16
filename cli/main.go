package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/recentralized/legal"
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
		in, out := pathsFor(path)
		fmt.Printf("Rendering %s to %s\n", path, out)
		store(out, render(in))
	}
}

func pathsFor(p string) (string, string) {
	base := path.Base(p)
	ext := path.Ext(p)
	name := strings.Replace(base, ext, "", 1)
	output := strings.Replace(base, ext, ".html", 1)
	return name, path.Join(dstDir, output)
}

func render(name string) []byte {
	output, err := legal.HTML(name, legal.DefaultVariables)
	if err != nil {
		fmt.Printf("Failed to render: %v\n", err)
		os.Exit(1)
	}
	return output
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
