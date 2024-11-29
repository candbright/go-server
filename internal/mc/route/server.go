package route

import (
	"context"
	"github.com/candbright/go-server/internal/mc/model"
	"github.com/candbright/go-server/pkg/rest"
	"github.com/gin-gonic/gin"
)

func init() {
	RegisterRoute(func(e *gin.Engine) {
		e.POST("/server/info/get", rest.H(getCurrentServerInfo))
		e.POST("/server/download_start", rest.H(startDownloadServer))
		e.POST("/server/download_status", rest.H(statusDownloadServer))
	})
}

type downloadStatusRsp struct {
	Downloading bool `json:"downloading"`
}

func getCurrentServerInfo(c *gin.Context, ctx context.Context) error {
	current := manager.CurrentServer()
	info := model.ServerInfo{
		Version: current.Version(),
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

	return rest.Json(info)
}

func startDownloadServer(c *gin.Context, ctx context.Context) error {
	return manager.CurrentServer().StartDownload()
}

func statusDownloadServer(c *gin.Context, ctx context.Context) error {
	return rest.Json(downloadStatusRsp{
		Downloading: manager.CurrentServer().Downloading(),
	})
}