package main

import (
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("static files", "./static")

	pkg, err := build.Default.Import("cmd", "", build.FindOnly)
	if err != nil {
		log.Fatal(err)
	}

	writeFile(pkg.Dir, "/trace/trace.go", box, "trace.go")
	writeFile(pkg.Dir, "/../../misc/trace/trace_viewer_full.html", box, "trace_viewer_full.html")
	writeFile(pkg.Dir, "/../../misc/trace/webcomponents.min.js", box, "webcomponents.min.js")

	log.Printf("running command: go install cmd/trace ...")
	c, err := exec.Command("go", "install", "cmd/trace").CombinedOutput()
	if err != nil {
		log.Fatal(string(c), err)
	}
}

func writeFile(dir, path string, box *packr.Box, name string) {
	b, err := box.Find(name)
	if err != nil {
		log.Fatal(err)
	}

	filePath := dir + filepath.FromSlash(path)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Printf("patching file %s ...\n", filePath)
	n, err := file.Write(b)
	if err != nil {
		log.Fatal(err)
	}

	if n != len(b) {
		log.Fatalf("write length: %d, should be: %d\n", n, len(b))
	}
}
