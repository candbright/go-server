package route

import (
	"net/http"

	"github.com/candbright/go-server/pkg/rest"
	"github.com/gin-gonic/gin"
)

func init() {
	registerRoute(func(e *gin.Engine) {
		e.POST("/server/current/allowlist/get", rest.H(getAllowList))
	})
}

func getAllowList(c *gin.Context) error {
	allowList, err := manager.CurrentServer().AllowList()
	if err != nil {
		return rest.ErrorWithStatus(http.StatusNotFound, err)
	}
	return rest.Json(allowList.GetAll())
}
