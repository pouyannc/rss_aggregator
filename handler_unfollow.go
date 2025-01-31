package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/pouyannc/rss_aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("expecting one argument (url)")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollowsRow(context.Background(), database.DeleteFeedFollowsRowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s unfollowed %s\n", user.Name, feed.Name)

	return nil
}
