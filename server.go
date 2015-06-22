package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
)

func loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("username")
	pw := request.FormValue("password")

	target := "/"
	if name != "" && pw != "" {
		target = "/user"
	}
	http.Redirect(response, request, target, 302)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("www/user/index.html")

	if err != nil {
		http.NotFound(w, r)
	} else {
		fmt.Fprintf(w, string(b))
	}
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/login/", loginHandler).Methods("POST")
	r.HandleFunc("/user/", userHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("www/")))
	http.Handle("/", r)

	err := http.ListenAndServe(":8082", r)
	fmt.Println(err.Error())
}
