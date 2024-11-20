package cmds

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jmsMaupin1/gator/internal/database"
)

func Browse(s *State, cmd Command, user database.User) error {
	var limit int32
	var err error

	if len(cmd.Args) < 1 {
		limit = 2
	} else {
		i, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("Error converting limit to int: %v", err)
		}

		limit = int32(i)
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: limit,
	})

	if err != nil {
		return fmt.Errorf("Browse error: %v", err)
	}

	for _, post := range posts {
		fmt.Printf("From: %s on %s\n", post.FeedName, post.PublishedAt.Format("Mon Jan 1"))
		fmt.Printf("========== %s ==========\n", post.Title)
		fmt.Println(post.Description)
		fmt.Println(post.Url)
		fmt.Println("========================================")
	}
	return nil
}
