package auth

import (
	"net/http"
	"strconv"

	"github.com/GuilhermeVendramini/golang-cms/core/modules/users"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

// cookie handling
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// GetUser get current user
func GetUser(r *http.Request) map[string]interface{} {
	user := make(map[string]interface{})
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			user["name"] = cookieValue["name"]
			user["email"] = cookieValue["email"]
			user["admin"] = cookieValue["admin"]
		}
	}
	return user
}

// SetSession cookie
func SetSession(user users.User, w http.ResponseWriter) {
	value := map[string]string{
		"name":  user.Name,
		"email": user.Email,
		"admin": strconv.FormatBool(user.Admin),
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

// ClearSession cookie
func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

// CheckPasswordHash check password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
