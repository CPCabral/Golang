package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

}

func BuscarUsuario