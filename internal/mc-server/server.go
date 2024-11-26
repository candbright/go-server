package server

import (
	"context"
	"github.com/candbright/go-log/log"
	"github.com/candbright/go-server/internal/mc-server/route"
	"github.com/candbright/go-server/pkg/config"
	"github.com/candbright/go-server/pkg/rest/handler"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	g sync.WaitGroup
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Serve() {
	ctx, cancel := context.WithCancel(context.Background())
	s.StartHTTPServer()
	s.StartWatcher(ctx)
	s.WaitSignal()
	cancel()
	s.g.Wait()
}

func (s *Server) StartWatcher(ctx context.Context) {
	log.Info("Start Watcher...")
	s.g.Add(1)
	go func() {
		defer s.g.Done()
		testTicker := time.NewTicker(time.Hour * 1)
		defer testTicker.Stop()
	LOOP:
		for {
			select {
			case <-testTicker.C:
				testLoop()
			case <-ctx.Done():
				break LOOP
			}
		}
		log.Warn("Exit Watcher!")
	}()
}

func (s *Server) StartHTTPServer() {
	log.Info("Start HTTP Server...")
	s.g.Add(1)
	go func() {
		defer s.g.Done()
		engine := gin.New()
		engine.Use(gin.BasicAuth(
			map[string]string{
				config.Global.Get("server.username"): config.Global.Get("server.password"),
			},
		))
		engine.Use(handler.LogHandler())
		engine.Use(gin.Recovery())
		route.Init()
		route.Incubate(engine)
		_ = engine.Run(":" + strconv.Itoa(config.Global.GetInt("server.port")))
		log.Warn("Exit HTTP server!")
	}()
}

func (s *Server) WaitSignal() {
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
