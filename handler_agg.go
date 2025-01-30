package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	url := s.config.RSSURL

	rssfeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Println(*rssfeed)

	return nil
}
