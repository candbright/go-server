package core

import (
	_ "embed"
	"github.com/candbright/go-ssh/ssh"
	"path"

	"github.com/candbright/go-server/internal/mc/core/model"
	"github.com/candbright/go-server/pkg/dw"
)

type AllowList struct {
	Session ssh.Session
	Version string
	rootDir string
	*dw.DataWriter[model.AllowList]
}

type AllowListConfig struct {
	Session ssh.Session
	Version string
	RootDir string
}

func NewAllowList(config AllowListConfig) *AllowList {
	al := &AllowList{
		Session: config.Session,
		Version: config.Version,
		rootDir: config.RootDir,
	}
	al.Init()
	return al
}

func (al *AllowList) FileName() string {
	return "allowlist.json"
}

func (al *AllowList) Init() {
	al.DataWriter = dw.Default[model.AllowList](al.Session, path.Join(al.rootDir, al.FileName()))
}

func (al *AllowList) GetAll() model.AllowList {
	return al.Data
}
