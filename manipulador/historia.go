package manipulador

import (
	"fmt"
	"net/http"
)

// Historia retorna pagina de com Historia
func Historia(w http.ResponseWriter, r *http.Request) {
	if err := ModelosHistoria.ExecuteTemplate(w, "historia.html", nil); err != nil {
		http.Error(w, "Houve um erro no carregamento da pagina", http.StatusInternalServerError)
		fmt.Println("[HISTORIA] Erro na execução do modelo:", err.Error())
	}
}
