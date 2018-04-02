package contact

import (
	"log"
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/config"
	"github.com/julienschmidt/httprouter"
	gomail "gopkg.in/gomail.v2"
)

// Contact form
func Contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := config.TPL.ExecuteTemplate(w, "contact.html", nil)
	HandleError(w, err)
}

// Process Email
func Process(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	name := r.FormValue("name")
	email := r.FormValue("email")
	subj := r.FormValue("subject")
	message := r.FormValue("message")

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", "to-test1@test.com", "to-test2@test.com")
	m.SetAddressHeader("Cc", "copy-test@test.com", "Name")
	m.SetHeader("Subject", subj)
	m.SetBody("text/html", "<b>Name:"+name+"</b><br>"+message)
	// m.Attach("/home/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 465, "test@gmail.com", "password")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	//TODO success or fail message
	http.Redirect(w, r, "/contact", http.StatusSeeOther)
}

// HandleError return Status Internal Server Error
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
