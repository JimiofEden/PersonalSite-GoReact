package main

import (
	"embed"
	"io/fs"
	"net/http"

	"main/config"
	"main/web"
)

var reactContent embed.FS


func main() {
	a := web.App{}

	// Get Configuration file settings
	config := config.GetConfig()

	// Try to start the server based on compiled code
	bin, err := fs.Sub(reactContent, "bin")
	if err != nil {
		panic(err)
	}
	var binFS = http.FS(bin)

	// Run the server
	a.Serve(binFS, config)
}