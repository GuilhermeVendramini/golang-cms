package article

import (
	"github.com/GuilhermeVendramini/golang-cms/core"
)

// Mux Article
func Mux() {
	core.Mux.GET("/articles", List)
	core.Mux.GET("/article/:url", Read)
	core.Mux.GET("/add/article", Add)
	core.Mux.POST("/add/article/process", AddProcess)
	core.Mux.GET("/delete/article/:url", Delete)
	core.Mux.POST("/delete/process/article/:url", DeleteProcess)
}
