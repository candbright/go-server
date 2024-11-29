package route

import (
	"context"
	"net/http"

	"github.com/candbright/go-server/pkg/rest"
	"github.com/gin-gonic/gin"
)

func init() {
	RegisterRoute(func(e *gin.Engine) {
		e.POST("/server/current/server_properties/get", rest.H(getServerProperties))
		e.POST("/server/current/server_properties/set", rest.H(setServerProperties))
	})
}

type setServerPropertiesReq map[string]string

func getServerProperties(c *gin.Context, ctx context.Context) error {
	serverProperties, err := manager.CurrentServer().ServerProperties()
	if err != nil {
		return rest.ErrorWithStatus(http.StatusNotFound, err)
	}
	return rest.Json(serverProperties.GetAll())
}

func setServerProperties(c *gin.Context, ctx context.Context) error {
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
