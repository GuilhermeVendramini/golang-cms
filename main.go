package main

import (
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/core"
	"github.com/GuilhermeVendramini/golang-cms/core/modules/users"
	"github.com/GuilhermeVendramini/golang-cms/modules/contact"
	"github.com/GuilhermeVendramini/golang-cms/modules/content/article"
)

func main() {
	// Core Users
	users.Mux()

	// Contact Mux
	contact.Mux()

	// Content Mux
	article.Mux()

	// Server Listen
	http.ListenAndServe(":8080", core.Mux)
}
