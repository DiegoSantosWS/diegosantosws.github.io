package manipulador

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/blog/model"
)

// Contato retorna pagina de com eventos
func Contato(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	msg := r.Form.Get("msg")

	pagina := model.Pagina{}
	pagina.Msg = msg
	if err := ModelosContato.ExecuteTemplate(w, "contato.html", pagina); err != nil {
		http.Error(w, "Houve um erro no carregamento da pagina", http.StatusInternalServerError)
		fmt.Println("[EVENTOS] Erro na execução do modelo:", err.Error())
	}
}

// EnviarContato retorna pagina de com eventos
func EnviarContato(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	nome := r.Form.Get("nome")
	telefone := r.Form.Get("telefone")
	to := r.Form.Get("email")
	texto := r.Form.Get("mensagem")

	msg := "Novo Contato\r\n"
	msg += fmt.Sprintf("Nome: %s", nome)
	msg += "\r\n"
	msg += fmt.Sprintf("Telefone: %s", telefone)
	msg += "\r\n"
	msg += fmt.Sprintf("Email: %s", to)
	msg += "\r\n"
	msg += "\r\n"
	msg += fmt.Sprintf("Mensagem:\r\n %s", texto)

	SendMail(msg, "E-mail Contato")
	http.Redirect(w, r, "/contato?msg=Agradecemos seu contato.#resp", http.StatusSeeOther)
}
