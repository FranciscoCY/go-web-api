package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Estructura
type Usuario struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

// Funciones
func Saludar(name string) string {
	return "Hola desde " + name + " una función"
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	funciones := template.FuncMap{
		"saludar": Saludar,
	}

	template := template.Must(template.New("index.html").
		Funcs(funciones).
		ParseFiles("index.html", "base.html"))

	c1 := Curso{"Go"}
	c2 := Curso{"Java"}
	c3 := Curso{"C#"}

	cursos := []Curso{c1, c2, c3}
	usuario := Usuario{"Francisco", 25, true, false, cursos}

	template.Execute(rw, usuario)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	server := &http.Server{
		Addr:    "localhost:2000",
		Handler: mux,
	}

	// Crear servidor
	fmt.Println("El servidor esta corriend en el puerto 200")
	fmt.Println("Run Server: http://localhost:2000/ ")
	log.Fatal(server.ListenAndServe())
}
