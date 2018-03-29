package users

import (
	"log"
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/julienschmidt/httprouter"
)

// User struct
type User struct {
	Name     string
	Email    string
	Password string
	Admin    bool
}

func users(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "users.html", nil)
	HandleError(w, err)
}

func addUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "user-add.html", nil)
	HandleError(w, err)
}

func addUserProcess(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error

	user := User{}
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	adm := false
	if r.FormValue("admin") == "on" {
		adm = true
	}

	user.Admin = adm

	currentEmail := r.FormValue("current-email")

	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Redirect(w, r, "/admin/add/user", http.StatusSeeOther)
	}

	if currentEmail != "" {
		_, err = Update(user, currentEmail)
	} else {
		_, err = Create(user)
	}

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

// HandleError return Status Internal Server Error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
