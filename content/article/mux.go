package article

import (
	"github.com/GuilhermeVendramini/golang-cms/core"
)

// Mux Article
func Mux() {
	core.Mux.GET("/articles", List)
	core.Mux.GET("/article/:url", Read)
	core.Mux.GET("/admin/add/article", Add)
	core.Mux.GET("/edit/article/:url", Edit)
	core.Mux.POST("/admin/add/article/process", ItemProcess)
	core.Mux.GET("/delete/article/:url", Delete)
	core.Mux.POST("/delete/process/article/:url", DeleteProcess)
	core.Mux.GET("/admin/content/article", AdminContentList)
}
