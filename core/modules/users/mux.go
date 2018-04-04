package users

import "github.com/GuilhermeVendramini/golang-cms/core/utils"

// Mux users
func Mux() {
	utils.Mux.GET("/user/:id", Read)
	utils.Mux.GET("/admin/users", List)
	utils.Mux.GET("/admin/add/user", Add)
	utils.Mux.POST("/admin/add/user/process", UserProcess)
	utils.Mux.GET("/admin/user/edit/:id", Edit)
	utils.Mux.GET("/admin/user/delete/:id", Delete)
	utils.Mux.POST("/admin/delete/process/user/:id", DeleteProcess)
	utils.Mux.GET("/login", Login)
	utils.Mux.POST("/login/process", LoginProcess)
	utils.Mux.GET("/logout", Logout)
}
