package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pouyannc/rss_aggregator/internal/config"
	"github.com/pouyannc/rss_aggregator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("failed to load database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := state{config: &cfg, db: dbQueries}

	cmds := commands{make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

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
