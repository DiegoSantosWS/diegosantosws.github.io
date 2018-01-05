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
	sobrenome := r.Form.Get("sobrenome")
	to := r.Form.Get("email")

	msg := "Novo Contato\r\n"
	msg += fmt.Sprintf("Nome: %s %s", nome, sobrenome)
	msg += "\r\n"
	msg += fmt.Sprintf("Email: %s", to)
	msg += "\r\n"
	msg += "Mensagem enviada via Go!"

	SendMail(msg)
	fmt.Printf("Olá %s %s seu email é: %s", nome, sobrenome, to)
	http.Redirect(w, r, "/contato?msg=Obrigado", http.StatusSeeOther)

}
