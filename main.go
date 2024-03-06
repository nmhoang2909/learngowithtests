package main

import (
	"log"
	"net/http"
	"os"

	"example.com/go-test/dj"
)

func myGreeterHanlder(w http.ResponseWriter, r *http.Request) {
	dj.Greet(w, "Hoang")
}

func main() {
	dj.Greet(os.Stdout, "Hoang")
	log.Fatal(http.ListenAndServe(":8989", http.HandlerFunc(myGreeterHanlder)))
}
