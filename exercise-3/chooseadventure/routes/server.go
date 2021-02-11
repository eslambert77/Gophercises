package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

type apiHandler struct{}

type arcHandler struct {
	arcs map[string]Arc
}

func (a arcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("chooseadventure", "routes", "html", "page.html")
	arcs := a.arcs
	urlFormatted := strings.Replace(r.URL.Path, "/", "", 1)
	if val, ok := arcs[urlFormatted]; ok {
		//Change this so it's not executing every time
		t := template.Must(template.New("chooseadventure").Funcs(template.FuncMap{"StringsJoin": strings.Join}).ParseFiles(filePath))
		err := t.Execute(w, val)
		if err != nil {
			panic(err)
		}
	}
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
	Options []Option `json:"options"`
}

type Option struct {
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
