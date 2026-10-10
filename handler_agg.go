package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time example: 1s, 1m, 1h>", cmd.Name)
	}

	timeBetweenReqsStr := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(timeBetweenReqsStr)
	if err != nil {
		return fmt.Errorf("invalid duration string: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
	ticket := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticket.C {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Printf("Error scraping feeds: %v\n", err)
		}
	}
}
