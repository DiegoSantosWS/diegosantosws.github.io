package manipulador

import (
	"html/template"
)

//Modelos de tamplates
var (
	Modelos           = template.Must(template.ParseFiles("html/tmplates/home.html"))
	ModelosContato    = template.Must(template.ParseFiles("html/tmplates/contato.html"))
	ModelosDepoimento = template.Must(template.ParseFiles("html/tmplates/depoimentos.html"))
	ModelosEventos    = template.Must(template.ParseFiles("html/tmplates/eventos.html"))
	ModelosFutebol    = template.Must(template.ParseFiles("html/tmplates/futebol.html"))
	ModelosHistoria   = template.Must(template.ParseFiles("html/tmplates/historia.html"))
)
