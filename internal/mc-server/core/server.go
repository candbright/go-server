package core

import (
	"embed"
	"fmt"
	"path"

	"github.com/candbright/go-ssh/ssh"
	"github.com/pkg/errors"
)

//go:embed template
var tmpl embed.FS

const (
	serviceDir = "/opt/bin"
)

type ServerConfig struct {
	Version string
	RootDir string
	Session ssh.Session
}

type Server struct {
	Version          string
	rootDir          string
	session          ssh.Session
	process          *Process
	allowList        *AllowList
	serverProperties *ServerProperties
}

func NewServer(cfg ServerConfig) (*Server, error) {

	server := &Server{
		Version: cfg.Version,
		rootDir: cfg.RootDir,
		session: cfg.Session,
	}
	process, err := NewProcess(ProcessConfig{
		Version: cfg.Version,
		RootDir: server.ServerDir(),
		Session: cfg.Session,
	})
	if err != nil {
		return nil, err
	}
	server.process = process
	exist, err := server.ServerExist()
	if err != nil {
		return nil, err
	}
	if !exist {
		return server, nil
	}
	err = server.Reload()
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (server *Server) DownloadUrl() string {
	return fmt.Sprintf("https://www.minecraft.net/bedrockdedicatedserver/bin-linux/bedrock-server-%s.zip", server.Version)
}

func (server *Server) ZipFileName() string {
	return fmt.Sprintf("bedrock-server-%s.zip", server.Version)
}

func (server *Server) ZipFile() string {
	return path.Join(server.rootDir, server.ZipFileName())
}

func (server *Server) ZipExist() (bool, error) {
	//zip文件是否存在
	return server.session.Exists(server.ZipFile())
}

func (server *Server) ServerDirName() string {
	return fmt.Sprintf("server-%s", server.Version)
}

func (server *Server) ServerDir() string {
	return path.Join(server.rootDir, server.ServerDirName())
}

func (server *Server) ServerExist() (bool, error) {
	//服务器目录是否存在
	return server.session.Exists(server.ServerDir())
}

func (server *Server) Active() (bool, error) {
	exist, err := server.ServerExist()
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	return server.process.Active(), nil
}

func (server *Server) ServerProperties() (*ServerProperties, error) {
	exist, err := server.ServerExist()
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("server not exist")
	}
	return server.serverProperties, nil
}

func (server *Server) AllowList() (*AllowList, error) {
	exist, err := server.ServerExist()
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("server not exist")
	}
	return server.allowList, nil
}

func (server *Server) Reload() error {
	server.serverProperties = NewServerProperties(ServerPropertiesConfig{
		Version: server.Version,
		RootDir: server.ServerDir(),
	})
	server.allowList = NewAllowList(AllowListConfig{
		Version: server.Version,
		RootDir: server.ServerDir(),
	})
	return nil
}

func (server *Server) Download() error {
	//检测是否存在当前版本的服务器目录
	existS, err := server.ServerExist()
	if err != nil {
		return err
	}
	if existS {
		return nil
	}
	//不存在当前版本的服务器目录，则检测是否存在当前版本的zip文件
	existZ, err := server.ZipExist()
	if err != nil {
		return err
	}
	//不存在当前版本的zip文件，先下载
	if !existZ {
		progressPath := path.Join(server.rootDir, fmt.Sprintf("progress-%s.log", server.Version))
		err = server.session.Run("wget", "--no-check-certificate", fmt.Sprintf("--timeout=%d", 600), "-o", progressPath, server.DownloadUrl(), "-P", server.rootDir)
		if err != nil {
			return err
		}
	}

	//解压zip文件
	err = server.session.MakeDirAll(server.ServerDir(), 0777)
	if err != nil {
		return err
	}
	err = server.session.Run("unzip", "-q", server.ZipFile(), "-d", server.ServerDir())
	if err != nil {
		return err
	}
	//4. reload
	err = server.Reload()
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) Delete() error {
	//如果服务器正在运行则先关闭
	active := server.process.Active()
	if active {
		err := server.process.Stop()
		if err != nil {
			return err
		}
	}
	//TODO: 备份

	//删除服务器目录
	existS, err := server.ServerExist()
	if err != nil {
		return err
	}
	if existS {
		err = server.session.RemoveAll(server.ServerDir())
		if err != nil {
			return err
		}
	}
	//删除zip文件
	existZ, err := server.ZipExist()
	if err != nil {
		return err
	}
	if existZ {
		err = server.session.Remove(server.ZipFile())
		if err != nil {
			return err
		}
	}
	return nil
}

func (server *Server) AllowListAdd(username string) error {
	//TODO
	return server.process.ExecCmd("")
}

func (server *Server) AllowListDelete(id string) error {
	//TODO
	return server.process.ExecCmd("")
}
