-- name: CreateFeedFollow :one
WITH inserted AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT users.name AS user_name, feeds.name AS feed_name FROM inserted
INNER JOIN users ON inserted.user_id = users.id
INNER JOIN feeds ON inserted.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feeds.user_id = users.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollowsRow :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;