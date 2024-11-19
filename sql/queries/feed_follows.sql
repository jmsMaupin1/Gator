-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    inserted_feed_follows.*,
    f.name AS feed_name,
    u.name AS user_name
FROM inserted_feed_follows
INNER JOIN feeds f on f.id = inserted_feed_follows.feed_id
INNER JOIN users u on u.id = inserted_feed_follows.user_id;

-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows;

-- name: GetFeedsFollowedByUser :many
SELECT 
    f.name as feed_name,
    u.name as user_name
FROM feed_follows
INNER JOIN feeds f on f.id = feed_follows.feed_id
INNER JOIN users u on u.id = feed_follows.user_id
WHERE feed_follows.user_id = $1;
