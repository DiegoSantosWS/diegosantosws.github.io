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
func SendMail(body, assunto string) {
	from := "tec.infor321@gmail.com"
	to := "tec.infor321@gmail.com"
	to1 := []string{"tec.infor321@gmail.com"}
	// Configurando as altenticações de e-mail
	auth := smtp.PlainAuth("", "...@dominio.com.br", "di@#12SD", "mail.host.com.br")
	// Conectar ao servidor, autenticar, configurar o remetente e o destinatario,
	// e envie o email tudo em uma unica etapa

	dados := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + assunto + " !\r\n" +
		body

	msg := []byte(dados)
	err := smtp.SendMail("mail.host.com.br:587", auth, "...@domino.com.br", to1, msg)
	if err != nil {
		log.Printf("[SENDMAIL] smtp error: %s", err.Error())
		return
	}
}
