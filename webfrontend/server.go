// +build server

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Page struct {
	html []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{html: body}, nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	title := "index"
	p, err := loadPage(title)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, string(p.html))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/login/"):]

	if len(title) == 0 {
		title = "login"
	}

	p, err := loadPage(title)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, string(p.html))
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/signup/"):]

	if len(title) == 0 {
		title = "signup"
	}

	p, err := loadPage(title)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, string(p.html))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login/", LoginHandler)
	r.HandleFunc("/signup/", SignupHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
