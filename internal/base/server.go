package base

import (
	"context"
	"github.com/candbright/go-log/log"
	"github.com/candbright/go-server/internal/spectrum/route"
	"github.com/candbright/go-server/pkg/config"
	"github.com/candbright/go-server/pkg/rest/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"sync"
	"syscall"
)

type Server struct {
	wg    sync.WaitGroup
	tasks []func(ctx context.Context)
}

func NewServer() *Server {
	return &Server{}
}

func NewHTTPServer() *Server {
	server := NewServer()
	server.AddTask(HTTPServer)
	return server
}

func ServeHTTP() {
	NewHTTPServer().Serve()
}

func (s *Server) AddTask(task func(ctx context.Context)) {
	if s.tasks == nil {
		s.tasks = make([]func(ctx context.Context), 0)
	}
	s.tasks = append(s.tasks, task)
}

func (s *Server) Serve() {
	if s.tasks == nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	for _, task := range s.tasks {
		funcName := reflect.TypeOf(task).Name()
		log.Infof("Start %s...", funcName)
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			task(ctx)
			log.Warnf("Exit %s!", funcName)
		}()
	}
	WaitSignal()
	cancel()
	s.wg.Wait()
}

func WaitSignal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		sig := <-signals
		if sig == syscall.SIGINT {
			log.Warn("Receive SIGINT: Force quit!")
			break
		} else if sig == syscall.SIGTERM {
			log.Warn("Receive SIGTERM: Force quit!")
			break
		}
	}
}

func HTTPServer(ctx context.Context) {
	engine := gin.New()
	engine.Use(gin.BasicAuth(
		map[string]string{
			config.Global.Get("server.username"): config.Global.Get("server.password"),
		},
	))
	engine.Use(handler.LogHandler())
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())
	route.Incubate(engine)
	_ = engine.Run(":" + strconv.Itoa(config.Global.GetInt("server.port")))
}
