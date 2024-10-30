package route

import (
	"github.com/candbright/go-ssh/ssh/options"
	"github.com/candbright/server-mc/internal/mc-server/config"
	"github.com/candbright/server-mc/internal/mc-server/core"
)

var servers *core.Servers

func Init() {
	servers = core.New(
		core.Config{
			Version: config.ServerConfig.Get("mc.version"),
			RootDir: config.ServerConfig.Get("mc.path"),
			SSHOptions: []options.Option{
				options.EnableSingle(),
				options.LocalHost(),
				//options.RemoteHostPWD(
				//	config.ServerConfig.Get("ssh.host"),
				//	uint16(config.ServerConfig.GetInt("ssh.port")),
				//	config.ServerConfig.Get("ssh.username"),
				//	config.ServerConfig.Get("ssh.password"),
				//),
				options.LogPath(config.ServerConfig.Get("ssh.log.path")),
			},
		},
	)
}
