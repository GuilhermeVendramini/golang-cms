package core

import (
	"log"
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "index.html", nil)
	HandleError(w, err)
}

func admin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "admin.html", nil)
	HandleError(w, err)
}

func content(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	q := r.URL.Query()
	tp := q.Get("type")

	vars := make(map[string]interface{})
	vars["Type"] = tp

	err := config.TPL.ExecuteTemplate(w, "content.html", vars)
	HandleError(w, err)
}

func login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
