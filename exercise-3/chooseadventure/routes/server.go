package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type apiHandler struct{}

type arcHandler struct {
	arcs map[string]Arc
}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := "exercise-3/chooseadventure/routes/html"
	if r.URL.Path == "/hello" {
		http.ServeFile(w, r, filePath+"hello.html")
	}
}

func (a arcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := "exercise-3/chooseadventure/routes/html"
	arcs := a.arcs
	var s string
	urlFormatted := strings.Replace("/", r.URL.Path, s, 1)
	if val, ok := arcs[urlFormatted]; ok {
		t, _ := template.ParseFiles(filePath + "page.html")
		t.Execute(w, val)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Henlo from root!")
}

func DefaultMux() *http.ServeMux {
	serv := http.NewServeMux()

	arcs := createArcs()
	serv.Handle("/intro", arcHandler{arcs})
	return serv
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []option `json:"options"`
}

type option struct {
	Text     string `json:"text"`
	StoryArc string `json:"arc"`
}

func parseJSON(b []byte) map[string]Arc {
	var ret map[string]Arc
	err := json.Unmarshal(b, &ret)
	if err != nil {
		panic("Can't parse the JSON")
	}

	return ret
}

func createArcs() map[string]Arc {
	jf, err := os.Open("gopher.json")
	defer jf.Close()
	if err != nil {
		panic("Can't open JSON file")
	}
	data, err := ioutil.ReadAll(jf)
	if err != nil {
		panic("Can't read JSON file")
	}
	arcs := parseJSON(data)
	return arcs
}
