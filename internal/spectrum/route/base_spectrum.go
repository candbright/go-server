package route

import (
	"github.com/candbright/go-server/internal/dao"
	"github.com/candbright/go-server/internal/spectrum/model"
	"github.com/gin-gonic/gin"
)

type CreateBaseSpectrumItem struct {
	FirstBoot int   `json:"first_boot"`
	Notes     []int `json:"notes"`
}

func init() {
	RegisterRoute(func(engine *gin.Engine) {
		engine.POST("/base_spectrum/scan", dao.Scan[model.BaseSpectrum])
		engine.POST("/base_spectrum/create", dao.Create[model.BaseSpectrum, CreateBaseSpectrumItem])
	})
}
