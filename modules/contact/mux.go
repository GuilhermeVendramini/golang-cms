package contact

import (
	"github.com/GuilhermeVendramini/golang-cms/core"
)

// Mux Article
func Mux() {
	core.Mux.GET("/contact", Contact)
}
