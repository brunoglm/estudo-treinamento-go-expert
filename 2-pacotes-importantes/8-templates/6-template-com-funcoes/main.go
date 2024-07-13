package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	curso := Curso{"golang 5", 50}
	t := template.New("CursoTemplate")
	t = t.Funcs(template.FuncMap{
		"toUpper": ToUpper,
	})
	t, err := t.Parse("Curso: {{.Nome | toUpper}} - Curso: {{toUpper .Nome}} - Carga Horária: {{.CargaHoraria}}")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
