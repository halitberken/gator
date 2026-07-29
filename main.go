package main

import (
	"log"
	"os"

	"github.com/halitberken/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	s := state{
		config: &cfg,
	}
	cmds := commands{
		commands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
		os.Exit(1)
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(1)
	}
}
