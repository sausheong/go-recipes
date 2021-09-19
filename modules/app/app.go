package app

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	vars := mux.Vars(r)
	t, _ := template.ParseFiles(dir + "/static/index.html")
	t.Execute(w, vars["category"])
}
