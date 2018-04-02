package contact

import (
	"log"
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/julienschmidt/httprouter"
)

// Contact form
func Contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "contact.html", nil)
	HandleError(w, err)
}

// HandleError return Status Internal Server Error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
