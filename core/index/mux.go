package index

import (
	"net/http"

	"github.com/GuilhermeVendramini/golang-cms/core/utils"
)

// Mux index
func Mux() {
	utils.Mux.ServeFiles("/static/*filepath", http.Dir("static"))
	utils.Mux.GET("/", index)
}
