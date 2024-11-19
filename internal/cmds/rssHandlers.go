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

	f, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: int32(uuid.New().ID()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
		Url: u.String(),
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	feed, err := s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: int32(uuid.New().ID()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: f.ID,
	})

	fmt.Println(fmt.Sprintf("Name: %s\nUser: %s", feed.FeedName, feed.UserName))
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

func FollowFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Expected url, no url was supplied")
	}

	ctx := context.Background()

	feed, err := s.DB.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil {
		return err
	}

	user, err := s.DB.GetUserByName(ctx, s.CFG.CurrentUserName)
	if err != nil {
		return err
	}

	uuid := uuid.New()

	feed_follows, err := s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: int32(uuid.ID()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}


	fmt.Println(fmt.Sprintf("Name: %s\nUser: %s", feed_follows.FeedName, feed_follows.UserName))

	return nil
}


func GetFollowsForCurrentUser(s *State, cmd Command) error {
	ctx := context.Background()

	user, err := s.DB.GetUserByName(ctx, s.CFG.CurrentUserName)
	if err != nil {
		return err
	}

	feed_follows, err := s.DB.GetFeedsFollowedByUser(ctx, user.ID)
	if err != nil {
		return err
	}

	fmt.Println(feed_follows)

	return nil
}

func ResetFeeds(s *State, cmd Command) error {
	ctx := context.Background()
	
	if err := s.DB.DeleteFeedFollows(ctx); err != nil {
		return err
	}

	if err := s.DB.DeleteFeeds(ctx); err != nil {
		return err
	}	

	return nil
}
