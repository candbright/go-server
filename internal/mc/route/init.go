package route

import (
	"github.com/candbright/go-server/internal/mc/core"
	"github.com/candbright/go-server/pkg/config"
	"github.com/candbright/go-ssh/ssh/options"
)

var manager *core.ServerManager

func Init() {
	manager = core.New(
		core.Config{
			RootDir: config.Global.Get("mc.path"),
			SSHOptions: []options.Option{
				options.HostPWD(
					config.Global.Get("ssh.host"),
					uint16(config.Global.GetInt("ssh.port")),
					config.Global.Get("ssh.username"),
					config.Global.Get("ssh.password"),
				),
				options.LogPath(config.Global.Get("ssh.log.path")),
				options.Linux(),
				options.Single(config.Global.GetBool("ssh.single")),
			},
		},
	)
}
