package main

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/blog/manipulador"
)

func main() {
	http.HandleFunc("/style.css", manipulador.CssPrincipal)
	http.HandleFunc("/ima1.jpg", manipulador.Imagens1)
	http.HandleFunc("/2.jpg", manipulador.Imagens2)
	http.HandleFunc("/3.jpg", manipulador.Imagens3)
	http.HandleFunc("/", manipulador.Home)
	http.HandleFunc("/historia", manipulador.Historia)
	http.HandleFunc("/eventos", manipulador.Eventos)
	http.HandleFunc("/futebol", manipulador.Futebol)
	http.HandleFunc("/contato", manipulador.Contato)
	http.HandleFunc("/contatoEnviar", manipulador.EnviarContato)
	http.HandleFunc("/depoimentos", manipulador.Funcao)
	fmt.Println("Escutando...")
	http.ListenAndServe(":8181", nil)
}
