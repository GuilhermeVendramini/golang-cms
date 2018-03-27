package core

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Mux *httprouter.Router
var Mux *httprouter.Router

func init() {
	Mux = httprouter.New()
	Mux.ServeFiles("/static/*filepath", http.Dir("static"))
	Mux.GET("/", index)
	Mux.GET("/login", login)
}
