package cmds

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/jmsMaupin1/gator/internal/database"
)

func Agg(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: agg <time>")
	}

	dur, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	
	fmt.Println(fmt.Sprintf("Collecting feed every %v", dur))

	ticker := time.NewTicker(dur)
	for ;; <-ticker.C {
		err := ScrapeFeeds(s)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func AddFeed(s *State, cmd Command, user database.User) error {
	ctx := context.Background()
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Usage: addfeed <feed_name> <feed_url>")
	}

	u, err := url.ParseRequestURI(cmd.Args[1])
	if err != nil {
		return err
	}

	f, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
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
		ID: uuid.New(),
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

func FollowFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: follow <feed url>")
	}

	ctx := context.Background()

	feed, err := s.DB.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil {
		return err
	}

	feed_follows, err := s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
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

func GetFollowsForCurrentUser(s *State, cmd Command, user database.User) error {
	ctx := context.Background()

	feed_follows, err := s.DB.GetFeedsFollowedByUser(ctx, user.ID)
	if err != nil {
		return err
	}

	fmt.Println(feed_follows)

	return nil
}

func DeleteFeedFollowRecord(s *State, cmd Command, user database.User) error {
	ctx := context.Background()

	feed, err := s.DB.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil {
		return err
	}

	if err := s.DB.DeleteFeedFollowRecord(ctx, database.DeleteFeedFollowRecordParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}); err != nil {
		return nil
	}

	return nil
}

func ResetFeeds(s *State, cmd Command) error {
	ctx := context.Background()
	
	if err := s.DB.ResetFeedFollows(ctx); err != nil {
		return err
	}

	if err := s.DB.DeleteFeeds(ctx); err != nil {
		return err
	}	

	return nil
}
