package main

import (
	"embed"
	"io/fs"
	"main/web"
	"main/config"
	"net/http"
)

var reactContent embed.FS

func main() {
	// Get Configuration file settings
	config := config.GetConfig()

	// Try to start the server based on compiled code
	bin, err := fs.Sub(reactContent, "bin")
	if err != nil {
		panic(err)
	}
	var binFS = http.FS(bin)

	// Run the server
	web.Serve(binFS, config)
}