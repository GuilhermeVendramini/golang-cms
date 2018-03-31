package users

import "github.com/GuilhermeVendramini/golang-cms/core"

// Mux users
func Mux() {
	core.Mux.GET("/user/:id", Read)
	core.Mux.GET("/admin/users", List)
	core.Mux.GET("/admin/add/user", Add)
	core.Mux.POST("/admin/add/user/process", UserProcess)
	core.Mux.GET("/admin/user/edit/:id", Edit)
	core.Mux.GET("/admin/user/delete/:id", Delete)
	core.Mux.POST("/admin/delete/process/user/:id", DeleteProcess)
}
