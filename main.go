package main

import (
	"fmt"

	"github.com/pouyannc/rss_aggregator/internal/config"
)

func main() {
	fmt.Println(config.Read())
}
