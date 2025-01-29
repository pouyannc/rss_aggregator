package main

import (
	"github.com/pouyannc/rss_aggregator/internal/config"
	"github.com/pouyannc/rss_aggregator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}
