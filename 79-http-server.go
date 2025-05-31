/**
http server built using standard net/http package

*/

package main

import (
	"fmt"
	"net/http"
)

func handleGet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {

	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/post", handlePost)

	http.ListenAndServe(":8090", nil)
}
