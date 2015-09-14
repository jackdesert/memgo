package main

import (
	"github.com/jackdesert/memgo/Godeps/_workspace/src/github.com/bmizerany/pat"
	"github.com/jackdesert/memgo/handler"
	"log"
	"net/http"
	"runtime"
)

func main() {

	// 4 CPUs on this machine.
	runtime.GOMAXPROCS(32)

	// "pat" allows params from url string
	mux := pat.New()

	mux.Post("/", http.HandlerFunc(handler.Post))
	mux.Get("/:key", http.HandlerFunc(handler.Get))

	http.Handle("/", mux)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)

}
