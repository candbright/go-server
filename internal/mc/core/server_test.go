package core

import (
	"github.com/candbright/go-ssh/ssh"
	"github.com/candbright/go-ssh/ssh/options"
	"path"
	"testing"
)

const testRootDir = "/opt/mc"
const testVersion = "1.21.43.01"

func testServer(t *testing.T) *Server {
	session, err := ssh.NewSession([]options.Option{
		options.EnableSingle(),
		options.LocalHost(),
		//options.RemoteHostPWD(
		//	config.ServerConfig.Get("ssh.host"),
		//	uint16(config.ServerConfig.GetInt("ssh.port")),
		//	config.ServerConfig.Get("ssh.username"),
		//	config.ServerConfig.Get("ssh.password"),
		//),
		options.LogPath(path.Join(testRootDir, "ssh.log")),
	}...)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	server, err := NewServer(ServerConfig{
		Version: testVersion,
		RootDir: testRootDir,
		Session: session,
	})
	return server
}
func TestServer_Download(t *testing.T) {
	server := testServer(t)
	err := server.Download()
	if err != nil {
		t.Fatalf("%+v", err)
	}
}
