package admin

import (
	"log"
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/GuilhermeVendramini/golang-cms/core/modules/users"
	"github.com/julienschmidt/httprouter"
)

// Admin page
func Admin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users.UserIsLogged(w, r)
	user := users.GetLoggedUser(r)
	vars := make(map[string]interface{})
	vars["LoggedUser"] = user
	err := config.TPL.ExecuteTemplate(w, "admin.html", vars)
	HandleError(w, err)
}

// Content page
func Content(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users.UserIsLogged(w, r)
	q := r.URL.Query()
	tp := q.Get("type")

	vars := make(map[string]interface{})
	lUser := users.GetLoggedUser(r)

	vars["LoggedUser"] = lUser
	vars["Type"] = tp

	err := config.TPL.ExecuteTemplate(w, "content.html", vars)
	HandleError(w, err)
}

// HandleError return Status Internal Server Error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
