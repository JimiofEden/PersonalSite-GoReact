package main

import (
	"embed"
	"io/fs"
	"main/web"
	"net/http"
)

var reactContent embed.FS

func main() {
	bin, err := fs.Sub(reactContent, "bin")
	if err != nil {
		panic(err)
	}
	var binFS = http.FS(bin)
	web.Serve(binFS, ":9876")
}