package route

import (
	"net/http"

	"github.com/candbright/server-mc/pkg/xgin"
	"github.com/gin-gonic/gin"
)

func init() {
	registerRoute(func(e *gin.Engine) {
		e.POST("/server/current/server_properties/get", xgin.H(getServerProperties))
		e.POST("/server/current/server_properties/set", xgin.H(setServerProperties))
	})
}

type setServerPropertiesReq map[string]string

func getServerProperties(c *gin.Context) error {
	serverProperties, err := manager.CurrentServer().ServerProperties()
	if err != nil {
		return xgin.ErrorWithStatus(http.StatusNotFound, err)
	}
	return xgin.Json(serverProperties.GetAll())
}

func setServerProperties(c *gin.Context) error {
	req := new(setServerPropertiesReq)
	if err := c.ShouldBindJSON(req); err != nil {
		return err
	}

	serverProperties, err := manager.CurrentServer().ServerProperties()
	if err != nil {
		return err
	}
	err = serverProperties.SetAll(*req)
	if err != nil {
		return err
	}
	return nil
}
