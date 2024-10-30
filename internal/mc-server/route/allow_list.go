package route

import (
	"net/http"

	"github.com/candbright/server-mc/pkg/xgin"
	"github.com/gin-gonic/gin"
)

func init() {
	registerRoute(func(e *gin.Engine) {
		e.POST("/server/current/allowlist/get", xgin.H(getAllowList))
	})
}

func getAllowList(c *gin.Context) error {
	allowList, err := servers.CurrentServer().AllowList()
	if err != nil {
		return xgin.ErrorWithStatus(http.StatusNotFound, err)
	}
	return xgin.Json(allowList.GetAll())
}
