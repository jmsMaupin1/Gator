package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmsMaupin1/gator/internal/cmds"
	"github.com/jmsMaupin1/gator/internal/config"
	"github.com/jmsMaupin1/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}	

	commands := cmds.Commands {
		Commands: map[string]func(*cmds.State, cmds.Command) error{},
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	
	var state cmds.State = cmds.State{
		DB: dbQueries,
		CFG: &cfg,
	}	

	commands.RegisterCommands("login", cmds.Login)
	commands.RegisterCommands("register", cmds.Register)
	commands.RegisterCommands("reset", cmds.Reset)
	commands.RegisterCommands("users", cmds.Users)
	commands.RegisterCommands("agg", cmds.Agg)

	userInput := os.Args

	if len(userInput) < 2 {
		fmt.Println("Error: Not enough arguments")
		os.Exit(1)
	}

	var cmd = cmds.Command{
		Name: userInput[1],
		Args: userInput[2:],
	}

	if err := commands.Run(&state, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
