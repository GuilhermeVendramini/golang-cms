package admin

import "github.com/GuilhermeVendramini/golang-cms/core"

// Mux admin
func Mux() {
	core.Mux.GET("/admin", Admin)
	core.Mux.GET("/admin/content", Content)
}
