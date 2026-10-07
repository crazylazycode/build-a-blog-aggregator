package main

import (
	"context"
	"fmt"
	"time"

	"github.com/crazytobeawake/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user (%s): %w", s.cfg.CurrentUserName, err)
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed Created successfully!")
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("ID:                %s\n", feed.ID)
	fmt.Printf("Name:            %s\n", feed.Name)
	fmt.Printf("URL:              %s\n", feed.Url)
	fmt.Printf("User ID:       %s\n", feed.UserID)
	fmt.Printf("Created At: %s\n", feed.CreatedAt)
	fmt.Printf("Updated At: %s\n", feed.UpdatedAt)
}
