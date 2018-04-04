package utils

import "github.com/julienschmidt/httprouter"

// Mux *httprouter.Router
var Mux *httprouter.Router

func init() {
	Mux = httprouter.New()
}
