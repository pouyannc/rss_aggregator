package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/pouyannc/rss_aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return errors.New("expecting up to one argument (number of posts)")
	}

	limit := 2
	var err error
	if len(cmd.args) == 1 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return err
	}

	fmt.Println("Showing most recent posts:")
	fmt.Println("====================================================================")
	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Printf("Published: %v\n", post.PublishedAt.Time.Format("Mon Jan 2"))
		fmt.Printf("Link: %v\n", post.Url)
		fmt.Println(post.Description.String)
		fmt.Println("====================================================================")
	}

	return nil
}
