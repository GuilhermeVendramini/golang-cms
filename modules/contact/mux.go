package contact

import "github.com/GuilhermeVendramini/golang-cms/core/utils"

// Mux Article
func Mux() {
	utils.Mux.GET("/contact", Contact)
	utils.Mux.POST("/contact/process", Process)
}
