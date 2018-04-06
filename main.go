package main

import (
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/core/index"
	"github.com/GuilhermeVendramini/golang-cms/core/modules/admin"
	"github.com/GuilhermeVendramini/golang-cms/core/modules/users"
	"github.com/GuilhermeVendramini/golang-cms/core/utils"
	"github.com/GuilhermeVendramini/golang-cms/modules/contact"
	"github.com/GuilhermeVendramini/golang-cms/modules/content/article"
)

func main() {
	// Index Mux
	index.Mux()

	// Users Mux
	users.Mux()

	// Admin Mux
	admin.Mux()

	// Contact Mux
	contact.Mux()

	// Content Mux
	article.Mux()

	// Uncomment the line below to generate a demo user
	// demo.User()

	// Server Listen
	http.ListenAndServe(":8080", utils.Mux)
}
