package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){ //http://localhost:5000/home
		w.Write([]byte("Olá Mundo!"))
	}

func usuarios(w http.ResponseWriter, r *http.Request){ //http://localhost:5000/usuarios
		w.Write([]byte("Carregar pagina de usuários"))
	}

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB
	// CLIENTE - SERVIDOR

	//ROTAS - maneira de identificar o tipo de mensagem
	// URI - identificador do rescurso
	// Metodo - GET(buscar dados), POST(cadastrar dados), PUT(atualização), DELETE

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil)) 

}