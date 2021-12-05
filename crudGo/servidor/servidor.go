package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type produto struct {
	ID    uint32  `json:"id"`
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
}

type materiaPrima struct {
	ID      uint32 `json:"id"`
	Nome    string `json:"nome"`
	Estoque uint32 `json:"estoque"`
}

type insumo struct {
	IdProduto      uint32 `json:"id_produto"`
	IdMateriaPrima uint32 `json:"id_materia_prima"`
	Quantidade     uint32 `json:"quantidade"`
}

// CriarProduto cria um produto no banco
func CriarProduto(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}
	var produto produto
	if erro = json.Unmarshal(corpoRequisicao, &produto); erro != nil {
		w.Write([]byte("Erro ao converter o produto para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao converter conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into produto (nome, valor) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(produto.Nome, produto.Valor)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	// STATUS CODES

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Produto criado com sucesso! Id: %d", idInserido)))

}

// CriarMateriaPrima cria uma matetiria prima no banco
func CriarMateriaPrima(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}
	var materiaPrima materiaPrima
	if erro = json.Unmarshal(corpoRequisicao, &materiaPrima); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao converter conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into materia_prima (nome, estoque) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(materiaPrima.Nome, materiaPrima.Estoque)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	// STATUS CODES

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Matéria prima criada com sucesso! Id: %d", idInserido)))

}

// CriarInsumo cria um insumo no banco
func CriarInsumo(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}
	var insumo insumo
	if erro = json.Unmarshal(corpoRequisicao, &insumo); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao converter conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into insumo (id_produto, id_materia_prima, quantidade) values (?, ?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(insumo.IdProduto, insumo.IdMateriaPrima, insumo.Quantidade)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	// STATUS CODES

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Insumo criado com sucesso! Id: %d", idInserido)))

}

// BuscarProduto traz todos os usuários salvos no banco de dados
func BuscarProduto(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from produto")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os produtos"))
		return
	}
	defer linhas.Close()

	var produtos []produto
	for linhas.Next() {
		var produto produto

		if erro := linhas.Scan(&produto.ID, &produto.Nome, &produto.Valor); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		produtos = append(produtos, produto)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(produtos); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

// BuscarProdutoEspecifico traz um produto específico salvo no banco de dados
func BuscarProdutoEspecifico(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from produto where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}
	defer linha.Close()

	var produto produto //usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&produto.ID, &produto.Nome, &produto.Valor); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(produto); erro != nil {
		w.Write([]byte("Erro ao converter o produto para JSON!"))
		return
	}
}

// AtualizarProduto altera os dados de um usuário no banco de dados
func AtualizarProduto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var produto produto
	if erro := json.Unmarshal(corpoRequisicao, &produto); erro != nil {
		w.Write([]byte("Erro ao converter o produto para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update produto set nome = ?, valor = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(produto.Nome, produto.Valor, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeletarProduto remove um usuário do banco de dados
func DeletarProduto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from produto where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o produto!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
