package cmds

import (
	"fmt"

	"github.com/jmsMaupin1/gator/internal/config"
	"github.com/jmsMaupin1/gator/internal/database"
)

type State struct {
	DB *database.Queries
	CFG *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Commands map[string]func(*State, Command) error
}

func (c *Commands) RegisterCommands(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}

func (c *Commands) Run(state *State, cmd Command) error {
	f, ok := c.Commands[cmd.Name]
	
	if !ok {
		return fmt.Errorf("command %s does not exist", cmd.Name)
	}

	if err := f(state, cmd); err != nil {
		return err
	}

	return nil
}

func Reset(s *State, cmd Command) error {
	if err := ResetFeeds(s, cmd); err != nil {
		return err
	}

	if err := ResetUsers(s, cmd); err != nil {
		return err
	}	

	return nil
}
