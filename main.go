package main

import (
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/core"
)

func main() {
	mux := core.Mux()
	http.ListenAndServe(":8080", mux)
}
