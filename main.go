package main

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/blog/manipulador"
)

func main() {
	//Routes to files statics
	//http.Handle("/static", http.FileServer(http.Dir("static")))
	http.HandleFunc("/style.css", manipulador.CSSPrincipal)
	http.HandleFunc("/ima1.jpg", manipulador.Imagens1)
	//http.HandleFunc("/2.jpg", manipulador.Imagens2)
	//http.HandleFunc("/3.jpg", manipulador.Imagens3)

	//Routers of pages
	http.HandleFunc("/", manipulador.Home)
	http.HandleFunc("/historia", manipulador.Historia)
	http.HandleFunc("/eventos", manipulador.Eventos)
	http.HandleFunc("/futebol", manipulador.Futebol)
	http.HandleFunc("/contato", manipulador.Contato)
	http.HandleFunc("/contatoEnviar", manipulador.EnviarContato)
	http.HandleFunc("/depoimentos", manipulador.Depoimento)
	http.HandleFunc("/enviarDepoimento", manipulador.EnviarDepoimento)
	fmt.Println("Escutando...")
	http.ListenAndServe(":8181", nil)
}
