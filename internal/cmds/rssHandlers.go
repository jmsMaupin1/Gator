package cmds

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/jmsMaupin1/gator/internal/database"
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

func AddFeed(s *State, cmd Command) error {
	ctx := context.Background()
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Not enough arguments expected name and url")
	}

	user, err := s.DB.GetUserByName(ctx, s.CFG.CurrentUserName)
	if err != nil {
		return err
	}

	u, err := url.ParseRequestURI(cmd.Args[1])
	if err != nil {
		return err
	}

	uuid := uuid.New()

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: int32(uuid.ID()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
		Url: u.String(),
		UserID: user.ID,
	})

	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func GetFeeds(s *State, _ Command) error {
	feeds, err := s.DB.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}
	return nil
}

func ResetFeeds(s *State, cmd Command) error {
	if err := s.DB.DeleteFeeds(context.Background()); err != nil {
		return err
	}

	return nil
}
