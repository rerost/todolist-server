package app

import (
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/rerost/todolist-server/app/server"
)

// Run starts the grapiserver.
func Run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewTaskServiceServer(),
		),
	)
	return s.Serve()
}
