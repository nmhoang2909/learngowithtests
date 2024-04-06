package main

import (
	"net/http"

	"example.com/go-test/dj"
)

func myGreeterHanlder(w http.ResponseWriter, r *http.Request) {
	dj.Greet(w, "Hoang")
}

func main() {
}
