package cmds

import (
	"context"
	"fmt"

	"github.com/jmsMaupin1/gator/internal/config"
	"github.com/jmsMaupin1/gator/internal/database"
	"github.com/jmsMaupin1/gator/internal/rss"
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

func ScrapeFeeds(s *State) error {
	ctx := context.Background()
	c := rss.NewClient()

	nextFeed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	if err := s.DB.MarkFeedFetched(ctx, nextFeed.ID); err != nil {
		return err
	}

	feed, err := c.FetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return err
	}

	fmt.Println(feed.Channel.Item)

	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
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
