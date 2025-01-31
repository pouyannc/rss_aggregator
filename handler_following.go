package main

import (
	"context"
	"fmt"

	"github.com/pouyannc/rss_aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("%s is currently following:\n", s.config.CurrentUsername)
	for _, feed := range following {
		fmt.Println(feed.FeedName)
	}

	return nil
}
