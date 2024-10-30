package route

import (
	"github.com/candbright/server-mc/internal/mc-server/model"
	"github.com/candbright/server-mc/pkg/xgin"
	"github.com/gin-gonic/gin"
)

func init() {
	registerRoute(func(e *gin.Engine) {
		e.GET("/server/current/info/get", xgin.H(getCurrentServerInfo))
	})
}

func getCurrentServerInfo(c *gin.Context) error {
	current := servers.CurrentServer()
	info := model.ServerInfo{
		Version: current.Version,
	}

	exist, err := current.ServerExist()
	if err != nil {
		return err
	}
	info.Exist = exist

	active, _ := current.Active()
	info.Active = active

	serverProperties, _ := current.ServerProperties()
	if serverProperties != nil {
		info.ServerProperties = serverProperties.GetAll()
	}

	allowList, _ := current.AllowList()
	if allowList != nil {
		info.AllowList = allowList.GetAll()
	}

	return xgin.Json(info)
}
