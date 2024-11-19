package cmds

import (
	"fmt"
	"context"	
	"github.com/jmsMaupin1/gator/internal/rss"
)

func Agg(_ *State, _ Command) error {
	feedURL := "https://www.wagslane.dev/index.xml"

	c := rss.NewClient()
	feed, err := c.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
