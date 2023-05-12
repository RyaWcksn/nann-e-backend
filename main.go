package main

import (
	"github.com/nann-e-backend/config"
	"github.com/nann-e-backend/server"
)

func main() {
	cfg := config.Cfg
	s := server.NewService(cfg)
	s.Start()
}
