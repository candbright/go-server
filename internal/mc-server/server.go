package server

import (
	"context"
	"github.com/candbright/go-core/rest/handler"
	"github.com/candbright/go-log/log"
	"github.com/candbright/server-mc/internal/mc-server/config"
	"github.com/candbright/server-mc/internal/mc-server/route"
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
	s.StartHTTPServer(ctx)
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
		backupTick := time.NewTicker(time.Hour * 1)
		defer backupTick.Stop()
	LOOP:
		for {
			select {
			case <-backupTick.C:
				backup()
			case <-ctx.Done():
				break LOOP
			}
		}
		log.Warn("Exit Watcher!")
	}()
}

func (s *Server) StartHTTPServer(ctx context.Context) {
	log.Info("Start HTTP Server...")
	s.g.Add(1)
	go func() {
		defer s.g.Done()
		engine := gin.New()
		engine.Use(gin.BasicAuth(
			map[string]string{
				config.ServerConfig.Get("server.username"): config.ServerConfig.Get("server.password"),
			},
		))
		engine.Use(handler.LogHandler())
		engine.Use(gin.Recovery())
		route.Init()
		route.Incubate(engine)
		_ = engine.Run(":" + strconv.Itoa(config.ServerConfig.GetInt("server.port")))
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
