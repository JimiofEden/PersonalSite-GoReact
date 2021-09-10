package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	err := http.ListenAndServe(":9876", nil)
	if (err != null) {
		panic("Error: " + err.Error())
	}
}