package route

import (
	"github.com/candbright/go-ssh/ssh/options"
	"github.com/candbright/server-mc/internal/mc-server/config"
	"github.com/candbright/server-mc/internal/mc-server/core"
)

var manager *core.ServerManager

func Init() {
	manager = core.New(
		core.Config{
			RootDir: config.ServerConfig.Get("mc.path"),
			SSHOptions: []options.Option{
				options.HostPWD(
					config.ServerConfig.Get("ssh.host"),
					uint16(config.ServerConfig.GetInt("ssh.port")),
					config.ServerConfig.Get("ssh.username"),
					config.ServerConfig.Get("ssh.password"),
				),
				options.LogPath(config.ServerConfig.Get("ssh.log.path")),
				options.Linux(),
				options.Single(config.ServerConfig.GetBool("server.single")),
			},
		},
	)
}
