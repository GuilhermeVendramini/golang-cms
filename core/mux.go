package core

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Mux base httprouter
func Mux() *httprouter.Router {
	mux := httprouter.New()
	mux.ServeFiles("/static/*filepath", http.Dir("static"))
	mux.GET("/", index)
	mux.GET("/login", login)
	return mux
}
