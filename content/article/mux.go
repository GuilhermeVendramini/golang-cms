package article

import (
	"github.com/GuilhermeVendramini/golang-cms/core"
)

// Mux Article
func Mux() {
	core.Mux.GET("/add/article", Add)
	core.Mux.POST("/add/article/process", AddProcess)
}
