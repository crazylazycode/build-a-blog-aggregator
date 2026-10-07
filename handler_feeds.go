package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		fmt.Printf("* Name:      %s\n", feed.Name)
		fmt.Printf("  URL:       %s\n", feed.Url)
		fmt.Printf("  UserName:  %s\n", feed.UserName)
	}

	return nil
}
