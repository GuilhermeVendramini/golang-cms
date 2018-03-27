package main

import (
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/content/article"
	"github.com/GuilhermeVendramini/golang-cms/core"
)

func main() {
	// Content Mux
	article.Mux()
	// Server Listen
	http.ListenAndServe(":8080", core.Mux)
}
