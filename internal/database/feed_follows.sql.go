// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING id, created_at, updated_at, user_id, feed_id
)
SELECT 
    inserted_feed_follows.id, inserted_feed_follows.created_at, inserted_feed_follows.updated_at, inserted_feed_follows.user_id, inserted_feed_follows.feed_id,
    f.name AS feed_name,
    u.name AS user_name
FROM inserted_feed_follows
INNER JOIN feeds f on f.id = inserted_feed_follows.feed_id
INNER JOIN users u on u.id = inserted_feed_follows.user_id
`

type CreateFeedFollowParams struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int32
	FeedID    int32
}

type CreateFeedFollowRow struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int32
	FeedID    int32
	FeedName  string
	UserName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
		&i.FeedName,
		&i.UserName,
	)
	return i, err
}

const deleteFeedFollows = `-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows
`

func (q *Queries) DeleteFeedFollows(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollows)
	return err
}

const getFeedsFollowedByUser = `-- name: GetFeedsFollowedByUser :many
SELECT 
    f.name as feed_name,
    u.name as user_name
FROM feed_follows
INNER JOIN feeds f on f.id = feed_follows.feed_id
INNER JOIN users u on u.id = feed_follows.user_id
WHERE feed_follows.user_id = $1
`

type GetFeedsFollowedByUserRow struct {
	FeedName string
	UserName string
}

func (q *Queries) GetFeedsFollowedByUser(ctx context.Context, userID int32) ([]GetFeedsFollowedByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedsFollowedByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedsFollowedByUserRow
	for rows.Next() {
		var i GetFeedsFollowedByUserRow
		if err := rows.Scan(&i.FeedName, &i.UserName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}