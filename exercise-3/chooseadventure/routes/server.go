package routes

import "net/http"

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		http.ServeFile(w, r, "hello.html")
	}
}

//DefaultMux Configures a default server with /hello as the initial endpoint
func DefaultMux() *http.ServeMux {
	serv := http.NewServeMux()
	serv.Handle("/hello", apiHandler{})
	return serv
}
