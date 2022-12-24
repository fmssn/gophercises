package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/fmssn/gophercises/ex3/structs"
)

func main() {
	jsonBytes, err := os.ReadFile("story.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully read story.json")

	var story structs.Story
	json.Unmarshal(jsonBytes, &story)

	mux := defaultMux()
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
	if err != nil {
		panic(err)
	}
}
