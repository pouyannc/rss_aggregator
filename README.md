# rss_aggregator

Command line interface for aggregating RSS feeds. View recent posts from feeds you have added or follow.
Register separate users, each having their own feeds following list.

Requires Postgres and Go installed to run.

### Install the gator CLI using
`go install github.com/pouyannc/rss_aggregator@latest`

Create a JSON config file in your home directory called ".gatorconfig.json".
Input the following fields into the config file:
`{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable","current_user_name":""}`

### To run the program:
`rss_aggregator <command> [arguments]`

### Example commands:
`rss_aggregator register username`
To register a new user

`rss_aggregator login username`
To login to a different registered user

`rss_aggregator users`
To view a list of all registered users

`rss_aggregator addfeed url`
To add a new RSS feed

`rss_aggregator feeds`
View all existing feeds

`rss_aggregator follow url`
Follow a feed that has been added

`rss_aggregator agg 5s`
To aggregate posts from your followed feeds, this should run in a separate terminal and will aggregate posts every 5 seconds. Enter a different interval in the same format if desired (e.g. 30s, 10m, 1h)

`rss_aggregator browse [limit]`
View the most recent posts from your feed, if limit is not provided it will default to showing 2 posts