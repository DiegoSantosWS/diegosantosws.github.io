package manipulador

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

//Funcao é um manipulador
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador usando funcao num pacote")
}

// SendMail envia um email
func SendMail(body string) {
	from := "...@gmail.com"
	to := "...@dominio.com.br,...@gmail.com"
	to1 := []string{"...@dominio.com.br", "...@gmail.com", "...@gmail.com"}
	// Configurando as altenticações de e-mail
	auth := smtp.PlainAuth("", "...@dominio.com.br", "password", "mail.host.com.br")
	// Conectar ao servidor, autenticar, configurar o remetente e o destinatario,
	// e envie o email tudo em uma unica etapa

	dados := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Novo Contato!\r\n" +
		body

	msg := []byte(dados)
	err := smtp.SendMail("mail.hotst.com.br:587", auth, "diego@dominio.com.br", to1, msg)
	if err != nil {
		log.Printf("[SENDMAIL] smtp error: %s", err.Error())
		return
	}
}
