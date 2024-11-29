package main

import (
	"flag"
	"fmt"
	"github.com/candbright/go-log/log"
	"github.com/candbright/go-log/options"
	"github.com/candbright/go-server/internal/dao"
	"github.com/candbright/go-server/internal/spectrum"
	"github.com/candbright/go-server/pkg/config"
	"github.com/sirupsen/logrus"
)

var _BUILD_ = ""

var (
	help       bool
	version    bool
	configFile string
)

func init() {
	flag.BoolVar(&help, "h", false, "print help")
	flag.BoolVar(&version, "v", false, "print version")
	flag.StringVar(&configFile, "c", "conf/config.yaml", "configuration file path")
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

	err := config.InitFromFile(configFile)
	if err != nil {
		fmt.Println("Load config failed! error: ", err)
		return
	}

	err = log.Init(
		options.Path(config.Global.Get("log.path")),
		options.Level(func() logrus.Level {
			level, err := logrus.ParseLevel(config.Global.Get("log.level"))
			if err != nil {
				return logrus.InfoLevel
			}
			return level
		}),
		options.Format(&logrus.JSONFormatter{}),
		options.GlobalField("app_name", config.Global.Get("server.name")),
	)
	if err != nil {
		panic(err)
	}
	err = dao.Init(config.Global.Get("dao.driver"))
	if err != nil {
		panic(err)
	}
	spectrum.Serve()
}
