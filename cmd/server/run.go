package main

import (
	"os"

	"github.com/rerost/todolist-server/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := app.Run()
	if err != nil {
		return 1
	}
	return 0
}
