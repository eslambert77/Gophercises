package routes

import "net/http"

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		http.ServeFile(w, r, "hello.html")
	}
}

func DefaultMux() *http.ServeMux {
	serv := http.NewServeMux()
	serv.Handle("/hello", apiHandler{})
	return serv
}
