package core

import (
	"log"
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "index.html", nil)
	HandleError(w, err)
}

func login(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "login.html", nil)
	HandleError(w, err)
}

// HandleError return Status Internal Server Error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
