package core

import (
	_ "embed"
	"path"

	"github.com/candbright/server-mc/internal/mc-server/core/model"
	"github.com/candbright/server-mc/pkg/dw"
)

type AllowList struct {
	Version string
	rootDir string
	*dw.DataWriter[model.AllowList]
}

type AllowListConfig struct {
	Version string
	RootDir string
}

func NewAllowList(config AllowListConfig) *AllowList {
	al := &AllowList{
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
	al.DataWriter = dw.Default[model.AllowList](path.Join(al.rootDir, al.FileName()))
}

func (al *AllowList) GetAll() model.AllowList {
	return al.Data
}
