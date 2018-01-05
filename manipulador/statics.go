package manipulador

import (
	"net/http"
)

// CssPrincipal carrega i arquivo, obs não está sendo usada
func CssPrincipal(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/style.css")
}

// Imagens1 carrega i arquivo, obs não está sendo usada
func Imagens1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/imgs/ima1.jpg")
}

// Imagens2 carrega i arquivo, obs não está sendo usada
func Imagens2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/imgs/2.jpg")
}

// Imagens3 carrega i arquivo, obs não está sendo usada
func Imagens3(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/imgs/3.jpg")
}
