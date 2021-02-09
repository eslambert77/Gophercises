package main

import (
	"net/http"

	ca "github.com/eslambert77/Gophercises/exercise-3/chooseadventure/routes"
)

func main() {
	//Let's get the hello world function and JSON working
	mux := ca.DefaultMux()
	http.ListenAndServe(":8080", mux)
}
