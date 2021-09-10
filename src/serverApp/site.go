package main

import (
	"embed"
	"io/fs"
	"main/web"
	"net/http"
)

var reactContent embed.FS

func main() {
	dist, error := fs.Sub(reactContent, "dist")
	if error != nil {
		panic("Error: " + error)
	}
	var distFs = http.FS(dist)
	web.Serve(distFS, ":9876")
}