package article

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/julienschmidt/httprouter"
)

// Article struct
type Article struct {
	Title   string
	Teaser  string
	Body    string
	Tags    string
	Author  string
	URL     string
	Changed time.Time
	Created time.Time
}

// List articles
func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	items, err := GetAll()
	if err != nil {
		panic(err)
	}
	err = config.TPL.ExecuteTemplate(w, "articles.html", items)
	HandleError(w, err)
}

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

	if item.Title == "" || item.Body == "" || item.URL == "" {
		http.Redirect(w, r, "/add/article", http.StatusSeeOther)
	}
	_, err := Create(item)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, item.URL, http.StatusSeeOther)
}

// Read a specific tutorial
func Read(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	URL := r.URL.Path
	item, err := Get(URL)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	err = config.TPL.ExecuteTemplate(w, "article.html", item)
	HandleError(w, err)
}

// Delete return delete-confirm,html
func Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	URL := r.URL.Path
	URL = strings.Replace(URL, "/delete", "", 1)
	item, err := Get(URL)
	if err != nil {
		panic(err)
	}
	err = config.TPL.ExecuteTemplate(w, "delete-confirm.html", item)
	HandleError(w, err)
}

// DeleteProcess delete action
func DeleteProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	URL := r.FormValue("item-url")
	err := Remove(URL)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
	HandleError(w, err)
}

// HandleError return Status Internal Server Error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
