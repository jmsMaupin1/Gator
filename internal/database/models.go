// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

type Feed struct {
	ID            int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Url           string
	UserID        int32
	LastFetchedAt sql.NullTime
}

type FeedFollow struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int32
	FeedID    int32
}

type User struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}
