package article

import "github.com/GuilhermeVendramini/golang-cms/core/utils"

// Mux Article
func Mux() {
	utils.Mux.GET("/articles", List)
	utils.Mux.GET("/article/:url", Read)
	utils.Mux.GET("/api/article/:id", ReadJSON)
	utils.Mux.GET("/admin/add/article", Add)
	utils.Mux.GET("/admin/edit/article/:id", Edit)
	utils.Mux.POST("/admin/add/article/process", ItemProcess)
	utils.Mux.GET("/admin/delete/article/:id", Delete)
	utils.Mux.POST("/admin/delete/process/article/:url", DeleteProcess)
	utils.Mux.GET("/admin/content/article", AdminContentList)
}
