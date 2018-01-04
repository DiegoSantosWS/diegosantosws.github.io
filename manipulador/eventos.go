package manipulador

import (
	"fmt"
	"net/http"
)

// Eventos retorna pagina de com eventos
func Eventos(w http.ResponseWriter, r *http.Request) {
	if err := ModelosEventos.ExecuteTemplate(w, "eventos.html", nil); err != nil {
		http.Error(w, "Houve um erro no carregamento da pagina", http.StatusInternalServerError)
		fmt.Println("[EVENTOS] Erro na execução do modelo:", err.Error())
	}
}
