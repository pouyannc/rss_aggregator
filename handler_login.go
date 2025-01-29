package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("expecting one argument")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}
