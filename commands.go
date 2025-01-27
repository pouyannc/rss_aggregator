package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commandHandlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandHandlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commandHandlers[cmd.name]
	if !ok {
		return errors.New("command does not exist")
	}
	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
