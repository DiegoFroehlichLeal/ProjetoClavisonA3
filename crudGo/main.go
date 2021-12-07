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

	router.HandleFunc("/produto", servidor.CriarProduto).Methods(http.MethodPost)     //cria o produto
	router.HandleFunc("/mprima", servidor.CriarMateriaPrima).Methods(http.MethodPost) //cria o produto
	router.HandleFunc("/insumo", servidor.CriarInsumo).Methods(http.MethodPost)       //cria o produto

	//Produtos
	router.HandleFunc("/produtos", servidor.BuscarProduto).Methods(http.MethodGet)
	router.HandleFunc("/produtos/{id}", servidor.BuscarProdutoEspecifico).Methods(http.MethodGet)
	router.HandleFunc("/produtos/{id}", servidor.AtualizarProduto).Methods(http.MethodPut)
	router.HandleFunc("/produtos/{id}", servidor.DeletarProduto).Methods(http.MethodDelete)
	//Materia Prima
	router.HandleFunc("/mprimas", servidor.BuscarMateriaPrima).Methods(http.MethodGet)
	router.HandleFunc("/mprimas/{id}", servidor.BuscarMateriaPrimaEspecifica).Methods(http.MethodGet)
	router.HandleFunc("/mprimas/{id}", servidor.AtualizarMateriaPrima).Methods(http.MethodPut)
	router.HandleFunc("/mprimas/{id}", servidor.DeletarMateriaPrima).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)

}
