package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DiegoSantosWS/blog/model"
)

//Home é manibulador da requisição a hora /
func Home(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := Modelos.ExecuteTemplate(w, "home.html", pagina); err != nil {
		http.Error(w, "Houve um erro no carregamento da pagina", http.StatusInternalServerError)
		fmt.Println("[HOME] Erro na execução do modelo:", err.Error())
	}
}
