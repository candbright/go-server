package spectrum

import (
	"github.com/candbright/go-server/internal/base"
)

func Serve() {
	server := base.NewHTTPServer()
	server.Serve()
}
