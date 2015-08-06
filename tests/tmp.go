package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Cookie)
	http.HandleFunc("/2", Cookie2)
	http.ListenAndServe(":8080", nil)
}

func Cookie(w http.ResponseWriter, r *http.Request) {
	expire := time.Now().AddDate(0, 0, 1)
	ck := &http.Cookie{
		Name:    "myCookie",
		Value:   "hello",
		Path:    "/",
		Domain:  "localhost",
		Expires: expire,
		MaxAge:  3600,
	}
	http.SetCookie(w, ck)

	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}
func Cookie2(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 3600,
	}
	w.Header().Set("Set_Cookie", ck.String())
	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}
