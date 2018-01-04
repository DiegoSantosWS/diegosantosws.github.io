package manipulador

import (
	"net/http"
)

// CssPrincipal carrega i arquivo, obs não está sendo usada
func CssPrincipal(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/css/style.css")
}
