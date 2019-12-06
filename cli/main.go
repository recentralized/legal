package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/recentralized/legal"
	"github.com/recentralized/legal/policies"
)

var (
	srcDir = policies.Path
	dstDir = "dist"
	port   = "3333"
)

func main() {
	var (
		s bool
		w bool
	)

	flag.BoolVar(&s, "s", false, "run a server to preview content")
	flag.BoolVar(&w, "w", true, "write documents to dist dir")
	flag.Parse()

	if s {
		serve()
		os.Exit(0)
	}
	if w {
		write()
		os.Exit(0)
	}
}

func serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(stylesheet))
		if r.URL.Path == "/" {
			indexPage(w)
			return
		}
		name := strings.TrimPrefix(r.URL.Path, "/")
		result, err := render(name)
		if err != nil {
			fmt.Fprintf(w, "Failed to render: %v\n", err)
			return
		}
		w.Write(result)
	})
	log.Printf("Serving on http://localhost:%s", port)
	http.ListenAndServe(":"+port, nil)
}

var stylesheet = `
<style>
body{ margin: 0 auto; width: 50%; }
</style>
`

func indexPage(w io.Writer) {
	paths, err := filepath.Glob(srcDir + "/**/*.md")
	if err != nil {
		fmt.Printf("Failed to list files: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(w, `<h1>Legal Docs Preview</h1>`)
	for _, path := range paths {
		name, _ := pathsFor(path)
		fmt.Fprintf(w, `<li><a href="/%s">%s</a></li>`, name, strings.Title(name))
	}
}

func write() {
	paths, err := filepath.Glob(srcDir + "/**/*.md")
	if err != nil {
		fmt.Printf("Failed to list files: %v\n", err)
		os.Exit(1)
	}

	for _, fpath := range paths {
		in, out := pathsFor(fpath)
		fmt.Printf("Rendering %s to %s\n", path.Base(fpath), out)
		result, err := render(in)
		if err != nil {
			fmt.Printf("Failed to render: %v\n", err)
			os.Exit(1)
		}
		store(out, result)
	}
}

func pathsFor(p string) (string, string) {
	subDir := strings.Replace(path.Dir(p), srcDir, "", 1)
	base := path.Base(p)
	ext := path.Ext(p)
	name := strings.Replace(base, ext, "", 1)
	output := strings.Replace(base, ext, ".html", 1)
	return path.Join(".", subDir, name), path.Join(dstDir, subDir, output)
}

func render(name string) ([]byte, error) {
	output, err := legal.HTML(name, legal.DefaultVariables)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func store(fpath string, data []byte) {
	if err := os.MkdirAll(path.Dir(fpath), 0755); err != nil {
		fmt.Printf("Failed to create dir: %v\n", err)
	}
	f, err := os.Create(fpath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		os.Exit(1)
	}

	if _, err := f.Write(data); err != nil {
		fmt.Printf("Failed to write file: %v\n", err)
		os.Exit(1)
	}
}
