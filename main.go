package main

import (
	"log"
	"os"

	"github.com/pouyannc/rss_aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	programState := state{config: &cfg}

	cmds := commands{make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("no command provided")
	}

	cmdName := args[1]
	cmdArgs := args[2:]
	cmd := command{name: cmdName, args: cmdArgs}
	err = cmds.run(&programState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
