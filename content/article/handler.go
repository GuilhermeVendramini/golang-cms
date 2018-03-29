package article

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/GuilhermeVendramini/golang-cms/core/functions"
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

// Add call article-add.html to add new article
func Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "article-add.html", nil)
	HandleError(w, err)
}

// Edit call article-add.html to edit a article
func Edit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	URL := r.URL.Path
	URL = strings.Replace(URL, "/edit", "", 1)
	item, err := Get(URL)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	err = config.TPL.ExecuteTemplate(w, "article-add.html", item)
	HandleError(w, err)
}

// ItemProcess add or edit article process
func ItemProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error

	item := Article{}
	item.Title = r.FormValue("title")
	item.Teaser = r.FormValue("teaser")
	item.Body = r.FormValue("body")
	item.Tags = r.FormValue("tags")
	item.Author = r.FormValue("author")
	item.URL = r.FormValue("url")
	item.Changed = time.Now()

	currentURL := r.FormValue("current-url")

	if item.Title == "" || item.Body == "" || item.URL == "" {
		http.Redirect(w, r, "/admin/add/article", http.StatusSeeOther)
	}

	if currentURL != "" {
		item.Created = functions.StringToTime(r.FormValue("created"))
		_, err = Update(item, currentURL)
	} else {
		item.Created = time.Now()
		_, err = Create(item)
	}

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

// AdminContentList admin article list
func AdminContentList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	items, err := GetAll()
	if err != nil {
		panic(err)
	}

	vars := make(map[string]interface{})
	vars["Type"] = "article"
	vars["Items"] = items

	err = config.TPL.ExecuteTemplate(w, "content.html", vars)
	HandleError(w, err)
}

// HandleError return Status Internal Server Error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
