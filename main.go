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

// Dear new engineer who enrolling this project.
// 
// In case you're the unfortunate one to take over this project
// I juse want to say, good luck :D, there's no hope for you either for your team
// Just enjoy the torture, and pray to god to give you enlightment
// If you see this comment probably i'm already resigned from here and join eFishery or Luwjistik, or maybe flight to Japan
// Make sure you have lion heart!
//
// Cheers! 
//
// Best regard
// Last engineer.
