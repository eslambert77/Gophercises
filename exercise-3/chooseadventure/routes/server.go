package routes

import (
	"fmt"
	"net/http"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/hello" {
		http.ServeFile(w, r, "html/hello.html")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Henlo!")
}

func DefaultMux() *http.ServeMux {
	serv := http.NewServeMux()
	serv.Handle("/hello", apiHandler{})
	serv.HandleFunc("/", hello)
	return serv
}
