package dw

import (
	"encoding/json"
	"encoding/xml"
	"github.com/candbright/go-ssh/ssh"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

func Default[T any](session ssh.Session, path string) *DataWriter[T] {
	return Json[T](session, path)
}

func Json[T any](session ssh.Session, path string) *DataWriter[T] {
	cfg := Config{
		Session:   session,
		Path:      path,
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
	}
	return New[T](cfg)
}

func Xml[T any](session ssh.Session, path string) *DataWriter[T] {
	cfg := Config{
		Session:   session,
		Path:      path,
		Marshal:   xml.Marshal,
		Unmarshal: xml.Unmarshal,
	}
	return New[T](cfg)
}

func Yaml[T any](session ssh.Session, path string) *DataWriter[T] {
	cfg := Config{
		Session:   session,
		Path:      path,
		Marshal:   yaml.Marshal,
		Unmarshal: yaml.Unmarshal,
	}
	return New[T](cfg)
}

func Toml[T any](session ssh.Session, path string) *DataWriter[T] {
	cfg := Config{
		Session:   session,
		Path:      path,
		Marshal:   toml.Marshal,
		Unmarshal: toml.Unmarshal,
	}
	return New[T](cfg)
}
