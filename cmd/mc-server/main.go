package main

import (
	"flag"
	"fmt"

	"github.com/candbright/go-log/log"
	"github.com/candbright/go-log/options"
	server "github.com/candbright/server-mc/internal/mc-server"
	"github.com/candbright/server-mc/internal/mc-server/config"
	"github.com/sirupsen/logrus"
)

var _BUILD_ = ""

var (
	help    bool
	version bool
)

func init() {
	flag.BoolVar(&help, "h", false, "print help")
	flag.BoolVar(&version, "v", false, "print version")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	if version {
		fmt.Println(_BUILD_)
		return
	}

	err := log.Init(
		options.Path(config.ServerConfig.Get("log.path")),
		options.Level(func() logrus.Level {
			level, err := logrus.ParseLevel(config.ServerConfig.Get("log.level"))
			if err != nil {
				return logrus.InfoLevel
			}
			return level
		}),
		options.Format(&logrus.JSONFormatter{}),
		options.GlobalField("app_name", config.ServerConfig.Get("server.name")),
	)
	if err != nil {
		panic(err)
	}
	server := server.NewServer()
	server.Serve()
}
