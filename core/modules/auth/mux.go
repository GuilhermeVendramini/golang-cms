package auth

import "github.com/GuilhermeVendramini/golang-cms/core"

// Mux users
func Mux() {
	core.Mux.GET("/logout", Logout)
	core.Mux.POST("/login/process", LoginProcess)
}
