package article

import (
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/julienschmidt/httprouter"
)

// Add call article-add.html
func Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "article-add.html", nil)
	HandleError(w, err)
}

// AddProcess add article process
func AddProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	item := Article{}
	item.Title = r.FormValue("title")
	item.Teaser = r.FormValue("teaser")
	item.Body = r.FormValue("body")
	item.Tags = r.FormValue("tags")
	item.Author = r.FormValue("author")
	item.URL = "/article/" + r.FormValue("url")
	item.Changed = time.Now()
	item.Created = time.Now()

	prop := reflect.ValueOf(&item).Elem()
	for i := 0; i < prop.NumField(); i++ {
		if prop.Interface() == "" {
			http.Redirect(w, r, "/add/article", http.StatusSeeOther)
		}
	}
	_, err := Create(item)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// HandleError return Status Internal Server Error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
