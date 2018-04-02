package article

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/GuilhermeVendramini/golang-cms/core/functions"
	"github.com/GuilhermeVendramini/golang-cms/core/modules/file"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// Article struct
type Article struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title   string
	Teaser  string
	Body    string
	Image   string
	Tags    string
	Author  string
	URL     string
	Changed time.Time
	Created time.Time
}

/* List articles with redis
func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	cRedis := redis.Client
	var items []Article

	result, err := cRedis.Get("articles").Bytes()
	if err == redis.Nil {
		fmt.Println("Without redis")
		items, err = GetAll()
		if err != nil {
			panic(err)
		}

		strItems, err := json.Marshal(items)
		if err != nil {
			panic(err)
		}

		err = cRedis.Set("articles", strItems, 0).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("With redis")
		err = json.Unmarshal(result, &items)
		if err != nil {
			panic(err)
		}
	}

	err = config.TPL.ExecuteTemplate(w, "articles.html", items)
	HandleError(w, err)
} */

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
	ID := strings.Replace(URL, "/admin/edit/article/", "", 1)
	item, err := GetbyID(ID)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	val := make(map[string]interface{})
	val["Content"] = item
	val["Type"] = "article"
	val["URL"] = strings.Replace(item.URL, "/article/", "", 1)
	err = config.TPL.ExecuteTemplate(w, "article-add.html", val)
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
	item.Changed = time.Now()

	URL := strings.Replace(r.FormValue("url"), "/article/", "", 1)
	item.URL = "/article/" + URL

	ID := r.FormValue("item-id")

	if item.Title == "" || item.Body == "" || item.URL == "" {
		http.Redirect(w, r, "/admin/add/article", http.StatusSeeOther)
	}

	img := r.FormValue("file")
	newImg := file.Upload(w, r, "file-upload", "static/images")

	if newImg != "" {
		file.Delete(img)
		item.Image = newImg
	} else {
		rFile := r.FormValue("file-remove")
		if rFile == "true" {
			file.Delete(img)
			item.Image = ""
		} else {
			item.Image = img
		}
	}

	if ID != "" {
		item.Created = functions.StringToTime(r.FormValue("created"))
		_, err = Update(item, ID)
	} else {
		item.Created = time.Now()
		_, err = Create(item)
	}

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, item.URL, http.StatusSeeOther)
}

// Read a specific article
func Read(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	URL := r.URL.Path
	item, err := GetbyURL(URL)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	vars := make(map[string]interface{})
	vars["Type"] = "article"
	vars["Content"] = item
	vars["BodyHTML"] = template.HTML(item.Body)

	err = config.TPL.ExecuteTemplate(w, "article.html", vars)
	HandleError(w, err)
}

// Delete return delete-content.html
func Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	URL := r.URL.Path
	ID := strings.Replace(URL, "/admin/delete/article/", "", 1)
	item, err := GetbyID(ID)
	if err != nil {
		panic(err)
	}

	vars := make(map[string]interface{})
	vars["Type"] = "article"
	vars["Content"] = item

	err = config.TPL.ExecuteTemplate(w, "delete-content.html", vars)
	HandleError(w, err)
}

// DeleteProcess delete action
func DeleteProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ID := r.FormValue("item-id")

	item, err := GetbyID(ID)
	if err != nil {
		panic(err)
	}

	file.Delete(item.Image)

	err = Remove(ID)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/admin/content/article", http.StatusSeeOther)
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
	vars["Content"] = items

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
