package cmds

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmsMaupin1/gator/internal/database"
)

func Login(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Expected username, no username supplied")
	}

	user, err := s.DB.GetUserByName(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("User %s does not exist", cmd.Args[0])
	}

	if err := s.CFG.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%s logged in!", user.Name))

	return nil
}

func Register(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Error: Expeted username, no username supplied")
	}
	
	uuid := uuid.New()

	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID: int32(uuid.ID()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	})

	if err != nil {
		return fmt.Errorf("Username already exists: %v", err)
	}

	fmt.Println(fmt.Sprintf("user: %v", user))

	return Login(s, cmd)
}

func Reset(s *State, _ Command) error {
	if err := s.DB.DeleteUsers(context.Background()); err != nil {
		return err
	}

	return nil
}

func Users(s *State, _ Command) error {
	users, err := s.DB.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.CFG.CurrentUserName {
			fmt.Println(fmt.Sprintf("%s (current)", user.Name))
		} else {
			fmt.Println(user.Name)
		}
	}

	return nil
}
