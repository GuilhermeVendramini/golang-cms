package admin

import "github.com/GuilhermeVendramini/golang-cms/core/utils"

// Mux admin
func Mux() {
	utils.Mux.GET("/admin", Admin)
	utils.Mux.GET("/admin/content", Content)
}
