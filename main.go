package main

import "net/http"

func main() {

	http.HandleFunc("/api", handleGetPlaylistLengths)
	http.ListenAndServe(":8000", nil)
}
