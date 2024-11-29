package core

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/candbright/go-server/internal/base/utils"
	"github.com/candbright/go-ssh/ssh"
	"html/template"
	"path"

	"github.com/candbright/go-server/pkg/dw"
	"github.com/magiconair/properties"
	"github.com/pkg/errors"
)

var serverPropertiesFilter = map[string][]string{
	"gamemode": {"survival", "creative", "adventure"},
}

type ServerProperties struct {
	Session ssh.Session
	Version string
	rootDir string
	*dw.DataWriter[map[string]string]
}

type ServerPropertiesConfig struct {
	Session ssh.Session
	Version string
	RootDir string
}

func NewServerProperties(config ServerPropertiesConfig) *ServerProperties {
	sp := &ServerProperties{
		Session: config.Session,
		Version: config.Version,
		rootDir: config.RootDir,
	}
	sp.Init()
	return sp
}

func (sp *ServerProperties) FileName() string {
	return "server.properties"
}

func (sp *ServerProperties) TemplateFileName() string {
	return path.Join("template", fmt.Sprintf("%s-%s", sp.Version, sp.FileName()))
}

func (sp *ServerProperties) Init() {
	sp.DataWriter = dw.New[map[string]string](dw.Config{
		Session: sp.Session,
		Path:    path.Join(sp.rootDir, sp.FileName()),
		Marshal: func(v any) ([]byte, error) {
			content, err := template.ParseFS(tmpl, sp.TemplateFileName())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			var result bytes.Buffer
			err = content.Execute(&result, v)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			return result.Bytes(), nil
		},
		Unmarshal: func(data []byte, v any) error {
			var err error
			p := properties.MustLoadString(string(data))
			if mapPtr, ok := v.(*map[string]string); ok {
				*mapPtr = p.Map()
				return nil
			}
			err = p.Decode(v)
			if err != nil {
				return errors.WithStack(err)
			}
			return nil
		},
	})
}

func (sp *ServerProperties) GetAll() map[string]string {
	return sp.Data
}

func (sp *ServerProperties) Get(key string) string {
	return sp.Data[key]
}

func (sp *ServerProperties) SetAll(data map[string]string) error {
	for k, v := range data {
		sp.Set(k, v, false)
	}
	return sp.Write()
}

func (sp *ServerProperties) Set(k, v string, write bool) error {
	_, kExist := sp.Data[k]
	if !kExist {
		return errors.Errorf("unsupported key [%s]", k)
	}
	filter, fExist := serverPropertiesFilter[v]
	if fExist && !utils.Contains(filter, v) {
		return errors.Errorf("unsupported value [%s]", v)
	}
	sp.Data[k] = v
	if write {
		return sp.Write()
	} else {
		return nil
	}
}

func (sp *ServerProperties) GetServerName() string {
	return sp.Get("server-name")
}
