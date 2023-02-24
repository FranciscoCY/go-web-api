package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El método es + " + r.Method)
	fmt.Fprintln(rw, "Hola Mundo")
}

func PaginaNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "La página no funca", http.StatusNotFound)
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(rw, "Hola, %s tu edad es %s !!", name, age)
}

func main() {

	mux := http.NewServeMux()
	// Router
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/hola", PaginaNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	// Crear servidor
	fmt.Println("El servidor esta corriend en el puerto 3000")
	fmt.Println("Run Server: http://localhost:3000/ ")
	log.Fatal(server.ListenAndServe())
}
