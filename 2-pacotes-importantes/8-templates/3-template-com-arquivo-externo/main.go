package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	cursos := Cursos{
		{"Nome1", 10},
		{"Nome2", 40},
	}
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}
}
