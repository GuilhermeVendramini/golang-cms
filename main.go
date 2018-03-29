package main

import (
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/content/article"
	"github.com/GuilhermeVendramini/golang-cms/core"
	"github.com/GuilhermeVendramini/golang-cms/core/modules/users"
)

func main() {
	// Core Users
	users.Mux()

	// Content Mux
	article.Mux()

	// Server Listen
	http.ListenAndServe(":8080", core.Mux)
}
