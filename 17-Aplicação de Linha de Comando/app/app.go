package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca ips e nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca o nome dos servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host:= c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips{
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context){
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //name server
	if erro != nil {
		log.Fatal(erro)		
	}

	for _, servidor := range servidores{
		fmt.Println(servidor.Host)
	}
}

// na linha de comando:
//go mod init linha-de-comando -> gera o go.mod

// go run main.go ip --host amazon.com.br
// go run main.go servidores --host amazon.com.br

// go build -> compila, gera o executável linha-de-comando
// PS G:\Meu Drive\UDEMY\Golang\17-Aplicação de Linha de Comando> ./linha-de-comando ip --host amazon.com.br