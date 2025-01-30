package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pouyannc/rss_aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	args := cmd.args
	if len(args) != 2 {
		return errors.New("expecting two arguments (name and url)")
	}

	name := args[0]
	url := args[1]

	currUsername := s.config.CurrentUsername
	user, err := s.db.GetUser(context.Background(), currUsername)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
