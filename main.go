package main

import (
	"github.com/lunai-monster/lunar-stack/internal/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		server.Module,
	).Run()
}
