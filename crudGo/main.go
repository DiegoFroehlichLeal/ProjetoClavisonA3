package main

import (
	"crud/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE

	// CREATE - POST
	// READ - GET
	// UPDATE - PUT
	// DELETE - DELETE
	templates = template.Must(template.ParseGlob("templates/*.html"))
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods(http.MethodGet)

	router.HandleFunc("/produto", servidor.CriarProduto).Methods(http.MethodPost)           //cria o produto
	router.HandleFunc("/materiaprima", servidor.CriarMateriaPrima).Methods(http.MethodPost) //cria o produto
	router.HandleFunc("/insumo", servidor.CriarInsumo).Methods(http.MethodPost)             //cria o produto

	/*router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)*/

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)

}
