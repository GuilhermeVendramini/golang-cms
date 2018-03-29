package users

import "github.com/GuilhermeVendramini/golang-cms/core"

// Mux users
func Mux() {
	core.Mux.GET("/admin/users", users)
	core.Mux.GET("/admin/add/user", addUser)
	core.Mux.POST("/admin/add/user/process", addUserProcess)
}
