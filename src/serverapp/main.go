package main

import (
	"embed"
	"io/fs"
	"net/http"

	"main/web"
	"main/config"
)

var reactContent embed.FS

type Cache struct {
	Cache  *redis.Client
}

func main() {
	c = Cache{}

	// Get Configuration file settings
	config := config.GetConfig()

	// Try to start the server based on compiled code
	bin, err := fs.Sub(reactContent, "bin")
	if err != nil {
		panic(err)
	}
	var binFS = http.FS(bin)

	// Run the server
	c.web.Serve(binFS, config)
}