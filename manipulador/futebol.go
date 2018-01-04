package manipulador

import (
	"fmt"
	"net/http"
)

// Futebol retorna pagina de com futebol
func Futebol(w http.ResponseWriter, r *http.Request) {
	if err := ModelosFutebol.ExecuteTemplate(w, "futebol.html", nil); err != nil {
		http.Error(w, "Houve um erro no carregamento da pagina", http.StatusInternalServerError)
		fmt.Println("[FUDETBOL] Erro na execução do modelo:", err.Error())
	}
}
