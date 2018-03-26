package main

import (
	"log"
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.ServeFiles("/static/*filepath", http.Dir("static"))
	mux.GET("/", index)
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "index.html", nil)
	HandleError(w, err)
}

// HandleError return Status Internal Server Error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
