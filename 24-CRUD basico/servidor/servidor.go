package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct{
	ID uint32 `json: "id"`
	Nome string `json: "nome"`
	Email string `json: email`
}
// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o retorno da requisição!"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario em struct!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return 
	}
	defer db.Close()

	//MySQL: INSERT INTO USUARIOS (nome, email) values ("nome", "email")
	//PREPARE STATEMENT	
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)") //evitar ataque ao banco de dados
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return 
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return 
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return 
	}

	// STATUS CODE
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}

// BuscarUsuarios traz todos os usuários salvos no banco de dados 
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
	}
	defer db.Close()

	//SELECT * FROM USUARIOS
	linhas, erro := db.Query("select *  from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuarios"))
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil{
			w.Write([]byte("Erro ao scanear usuário"))
			return
		}
		// 1    joão    email
		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
			w.Write([]byte("Erro ao converter usuarios para JSON"))
			return
	}
}

//BuscarUsuario traz um usuario especifico salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32) 
	if erro != nil {
			w.Write([]byte("Erro ao coverter parametro para inteiro"))
			return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("erro ao conectar com o banco de dados"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil{
			w.Write([]byte("Erro ao scanear o usuario"))
			return
		}
	}
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil{
		w.Write([]byte("Erro ao converter o usuario para JSON"))
		return
	}
}

// AtualizarUsuario altera os dados de um usuario no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}
	var usuario usuario
	if erro:= json.Unmarshal(corpoRequisicao, &usuario); erro != nil{
		w.Write([]byte("Erro ao converter usuário para struct"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	// statement para qualquer operação que não seja consulta
	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o Usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeletarUsuario remove um usuario do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil{
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
		if erro != nil {
			w.Write([]byte("Erro ao criar o statement"))
			return
		}
		defer statement.Close()

		if _, erro := statement.Exec(ID); erro != nil{
			w.Write([]byte("Erro ao deletar o usuario"))
			return
		}

		w.WriteHeader(http.StatusNoContent)
}