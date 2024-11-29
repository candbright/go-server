package mc

import (
	"context"
	"github.com/candbright/go-server/internal/base"
	"github.com/candbright/go-server/internal/mc/route"
	"time"
)

func Serve() {
	route.Init()
	server := base.NewHTTPServer()
	server.AddTask(Watcher)
	server.Serve()
}

func Watcher(ctx context.Context) {
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
}

func testLoop() {

}
