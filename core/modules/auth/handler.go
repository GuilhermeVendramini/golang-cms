package auth

import (
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/core/modules/users"
	"github.com/julienschmidt/httprouter"
)

// LoginProcess process user register
func LoginProcess(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	email := req.FormValue("email")
	pass := req.FormValue("password")
	redirect := "/login"
	if email != "" && pass != "" {
		user := users.User{}
		//err := users.User.Find(bson.M{"name": name}).One(&user)
		_, err := users.GetbyEmail(email)
		if err != nil {
			http.Redirect(w, req, redirect, http.StatusSeeOther)
		}
		match := CheckPasswordHash(pass, user.Password)
		if match == false {
			http.Redirect(w, req, redirect, http.StatusSeeOther)
		}
		SetSession(user, w)
		redirect = "/admin"
	}
	http.Redirect(w, req, redirect, 302)
}

// Logout user
func Logout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ClearSession(w)
	http.Redirect(w, req, "/", 302)
}
