package main

import (
	"crud/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE

	// CREATE - POST
	// READ - GET
	// UPDATE - PUT
	// DELETE - DELETE
	router := mux.NewRouter()

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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://*:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	handler := c.Handler(router)

	fmt.Println("Ouvindo na porta 5000 11111")
	log.Fatal(http.ListenAndServe(":5000", handler))

}
