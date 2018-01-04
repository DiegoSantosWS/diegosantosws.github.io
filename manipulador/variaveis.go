package manipulador

import (
	"html/template"
)

//Modelos de tamplates
var (
	Modelos           = template.Must(template.ParseFiles("view/home.html"))
	ModelosContato    = template.Must(template.ParseFiles("view/contato.html"))
	ModelosDepoimento = template.Must(template.ParseFiles("view/depoimentos.html"))
	ModelosEventos    = template.Must(template.ParseFiles("view/eventos.html"))
	ModelosFutebol    = template.Must(template.ParseFiles("view/futebol.html"))
	ModelosHistoria   = template.Must(template.ParseFiles("view/historia.html"))
)
