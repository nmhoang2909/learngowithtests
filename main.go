package main

import (
	"net/http"
	"os"

	"example.com/go-test/dj"
	"example.com/go-test/mocking"
)

func myGreeterHanlder(w http.ResponseWriter, r *http.Request) {
	dj.Greet(w, "Hoang")
}

func main() {
	// dj.Greet(os.Stdout, "Hoang")
	// log.Fatal(http.ListenAndServe(":8989", http.HandlerFunc(myGreeterHanlder)))

	defaultSleeper := &mocking.DefaultSleeper{}
	mocking.CountDown(os.Stdout, defaultSleeper)
}
