package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Feeds:")
	fmt.Println("==========================================")

	for _, feed := range feeds {
		createdBy, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("Name: %v\n", feed.Name)
		fmt.Printf("URL: %v\n", feed.Url)
		fmt.Printf("Created by: %v\n", createdBy.Name)
		fmt.Println("==========================================")
	}

	return nil
}
