package users

import (
	"log"
	"net/http"
	"strings"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// User struct
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string
	Email    string
	Password string
	Admin    bool
}

// List all users
func List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := GetAll()
	if err != nil {
		panic(err)
	}
	err = config.TPL.ExecuteTemplate(w, "users.html", users)
	HandleError(w, err)
}

// Read a specific user
func Read(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	URL := r.URL.Path
	ID := strings.Replace(URL, "/user/", "", 1)
	user, err := GetbyID(ID)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	vars := make(map[string]interface{})
	vars["User"] = user

	err = config.TPL.ExecuteTemplate(w, "user.html", vars)
	HandleError(w, err)
}

// Add a new user
func Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "user-add.html", nil)
	HandleError(w, err)
}

// Edit call user-add.html to edit a user
func Edit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	URL := r.URL.Path
	ID := strings.Replace(URL, "/admin/user/edit/", "", 1)
	user, err := GetbyID(ID)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	val := make(map[string]interface{})
	val["User"] = user

	if user.Admin {
		val["IsAdmin"] = "checked"
	}

	err = config.TPL.ExecuteTemplate(w, "user-add.html", val)
	HandleError(w, err)
}

// UserProcess add or edit user
func UserProcess(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error

	user := User{}
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")

	user.Password, err = HashPassword(r.FormValue("password"))
	if err != nil {
		panic(err)
	}

	adm := false
	if r.FormValue("admin") == "on" {
		adm = true
	}

	user.Admin = adm

	ID := r.FormValue("user-id")

	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Redirect(w, r, "/admin/add/user", http.StatusSeeOther)
	}

	if ID != "" {
		_, err = Update(user, ID)
	} else {
		_, err = Create(user)
	}

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

// Delete return delete-user.html
func Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	URL := r.URL.Path
	ID := strings.Replace(URL, "/admin/user/delete/", "", 1)
	user, err := GetbyID(ID)
	if err != nil {
		panic(err)
	}

	vars := make(map[string]interface{})
	vars["User"] = user

	err = config.TPL.ExecuteTemplate(w, "delete-user.html", vars)
	HandleError(w, err)
}

// DeleteProcess delete action
func DeleteProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ID := r.FormValue("user-id")
	err := Remove(ID)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	HandleError(w, err)
}

// HashPassword Hash user password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// HandleError return Status Internal Server Error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
