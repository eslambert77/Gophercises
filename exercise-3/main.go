package main

import (
	"net/http"

	"github.com/eslambert77/Gophercises/exercise-3/chooseadventure/routes"
)

func main() {
	//Let's get the hello world function and JSON working
	serv := routes.DefaultMux()

	http.ListenAndServe(":8080", serv)

}
