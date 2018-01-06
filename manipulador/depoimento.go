package manipulador

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/blog/model"
)

//Depoimento retorna pagina do formulario de depoimento
func Depoimento(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	msg := r.Form.Get("msg")

	pagina := model.Pagina{}
	pagina.Msg = msg
	if err := ModelosDepoimento.ExecuteTemplate(w, "depoimentos.html", pagina); err != nil {
		http.Error(w, "Houve um erro no carregamento da pagina", http.StatusInternalServerError)
		fmt.Println("[DEPOIMENTOS] Erro na execução do modelo:", err.Error())
	}
}

//EnviarDepoimento dispara envio do form para email
func EnviarDepoimento(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	nome := r.Form.Get("nome")
	texto := r.Form.Get("mensagem")

	msg := "Novo Depoimento\r\n"
	msg += fmt.Sprintf("Nome: %s", nome)
	msg += "\r\n"
	msg += fmt.Sprintf("Mensagem:\r\n %s", texto)

	SendMail(msg, "Depoimento")
	http.Redirect(w, r, "/depoimentos?msg=Agradecemos seu depoimento.#resp", http.StatusSeeOther)
}
