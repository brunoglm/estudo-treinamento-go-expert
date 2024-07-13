package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	cursos := Cursos{
		{"Nome1", 10},
		{"Nome2", 20},
	}

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
		"template.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles(templates...))
		err := t.Execute(w, cursos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8282", nil)
}
