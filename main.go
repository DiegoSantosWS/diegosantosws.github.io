package main

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/blog/manipulador"
)

func main() {
	http.HandleFunc("/style.css", manipulador.CssPrincipal)
	http.HandleFunc("/", manipulador.Home)
	http.HandleFunc("/historia", manipulador.Historia)
	http.HandleFunc("/eventos", manipulador.Eventos)
	http.HandleFunc("/futebol", manipulador.Futebol)
	http.HandleFunc("/contato", manipulador.Funcao)
	http.HandleFunc("/depoimentos", manipulador.Funcao)
	fmt.Println("Escutando...")
	http.ListenAndServe(":8181", nil)
}
