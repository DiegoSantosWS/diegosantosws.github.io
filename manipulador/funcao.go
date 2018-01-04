package manipulador

import (
	"fmt"
	"net/http"
)

//Funcao é um manipulador
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador usando funcao num pacote")
}
