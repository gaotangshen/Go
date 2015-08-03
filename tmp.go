package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))

	err1 := http.ListenAndServe(":8080", mux)
	if err1 != nil {
		log.Fatal(err)
	}
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "url: "+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello. world")
}
