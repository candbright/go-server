package core

import (
	"path"

	"github.com/candbright/go-ssh/ssh"
	"github.com/candbright/go-ssh/ssh/options"
	"github.com/go-resty/resty/v2"
)

type Config struct {
	Version    string
	RootDir    string
	SSHOptions []options.Option
}

func New(cfg Config) *Servers {
	session, err := ssh.NewSession(cfg.SSHOptions...)
	if err != nil {
		panic(err)
	}
	server, err := NewServer(ServerConfig{
		Version: cfg.Version,
		RootDir: cfg.RootDir,
		Session: session,
	})
	if err != nil {
		panic(err)
	}
	servers := &Servers{
		session: session,
		current: server,
		client:  resty.New(),
	}
	return servers
}

type Servers struct {
	session  ssh.Session
	current  *Server
	previous map[string]*Server
	client   *resty.Client
}

func (servers *Servers) LatestVersion() (string, error) {
	//TODO
	/*resp, err := m.client.R().
		Get("https://www.minecraft.net/en-us/download/server/bedrock")
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(resp.Body()), nil*/
	return "1.20.62.02", nil
}

func (servers *Servers) VersionScan() ([]string, error) {
	//TODO
	return nil, nil
}

func (servers *Servers) CurrentServer() *Server {
	return servers.current
}

func (servers *Servers) Upgrade() error {
	//1. 获取最新版本
	oldServer := servers.current
	newVersion, err := servers.LatestVersion()
	if err != nil {
		return err
	}
	//2. 若最新版本和当前版本不同，则下载最新版本
	if newVersion == oldServer.Version {
		return nil
	}
	newServer, err := NewServer(ServerConfig{
		Version: newVersion,
		RootDir: oldServer.rootDir,
		Session: servers.session,
	})
	if err != nil {
		return err
	}
	err = newServer.Download()
	if err != nil {
		return err
	}
	servers.current = newServer
	servers.previous[oldServer.Version] = oldServer
	//3. 复制旧版本数据文件到新版本
	err = servers.session.Run("cp", "-r",
		path.Join(oldServer.ServerDir(), "world"),
		path.Join(servers.current.ServerDir()+"/"))
	if err != nil {
		return err
	}
	err = servers.session.Run("cp",
		path.Join(oldServer.ServerDir(), oldServer.allowList.FileName()),
		path.Join(servers.current.ServerDir(), servers.current.allowList.FileName()))
	if err != nil {
		return err
	}
	err = servers.session.Run("cp",
		path.Join(oldServer.ServerDir(), oldServer.serverProperties.FileName()),
		path.Join(servers.current.ServerDir(), servers.current.serverProperties.FileName()))
	if err != nil {
		return err
	}
	return nil
}
