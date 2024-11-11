package core

import (
	"path"

	"github.com/candbright/go-ssh/ssh"
	"github.com/candbright/go-ssh/ssh/options"
	"github.com/go-resty/resty/v2"
)

type Config struct {
	RootDir    string
	SSHOptions []options.Option
}

func New(cfg Config) *ServerManager {
	session, err := ssh.NewSession(cfg.SSHOptions...)
	if err != nil {
		panic(err)
	}
	manager := &ServerManager{
		session: session,
		client:  resty.New(),
	}
	latestVersion, err := manager.CurrentLatestVersion()
	if err != nil {
		panic(err)
	}
	server, err := NewServer(ServerConfig{
		Version: latestVersion,
		RootDir: cfg.RootDir,
		Session: session,
	})
	if err != nil {
		panic(err)
	}
	manager.current = server
	return manager
}

type ServerManager struct {
	session  ssh.Session
	current  *Server
	versions []string
	client   *resty.Client
}

func (manager *ServerManager) LatestVersion() (string, error) {
	//TODO
	/*resp, err := m.client.R().
		Get("https://www.minecraft.net/en-us/download/server/bedrock")
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(resp.Body()), nil*/
	return "1.21.43.01", nil
}

func (manager *ServerManager) CurrentLatestVersion() (string, error) {
	//TODO
	/*resp, err := m.client.R().
		Get("https://www.minecraft.net/en-us/download/server/bedrock")
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(resp.Body()), nil*/
	return "1.21.43.01", nil
}

func (manager *ServerManager) VersionScan() ([]string, error) {
	//TODO
	return nil, nil
}

func (manager *ServerManager) CurrentServer() *Server {
	return manager.current
}

func (manager *ServerManager) Upgrade() error {
	//1. 获取最新版本
	oldServer := manager.current
	newVersion, err := manager.LatestVersion()
	if err != nil {
		return err
	}
	//2. 若最新版本和当前版本不同，则下载最新版本
	if newVersion == oldServer.version {
		return nil
	}
	newServer, err := NewServer(ServerConfig{
		Version: newVersion,
		RootDir: oldServer.rootDir,
		Session: manager.session,
	})
	if err != nil {
		return err
	}
	err = newServer.Download()
	if err != nil {
		return err
	}
	manager.current = newServer
	//3. 复制旧版本数据文件到新版本
	err = manager.session.Run("cp", "-r",
		path.Join(oldServer.ServerDir(), "world"),
		path.Join(manager.current.ServerDir()+"/"))
	if err != nil {
		return err
	}
	err = manager.session.Run("cp",
		path.Join(oldServer.ServerDir(), oldServer.allowList.FileName()),
		path.Join(manager.current.ServerDir(), manager.current.allowList.FileName()))
	if err != nil {
		return err
	}
	err = manager.session.Run("cp",
		path.Join(oldServer.ServerDir(), oldServer.serverProperties.FileName()),
		path.Join(manager.current.ServerDir(), manager.current.serverProperties.FileName()))
	if err != nil {
		return err
	}
	return nil
}
