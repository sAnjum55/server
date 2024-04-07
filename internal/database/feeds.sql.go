// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
insert into feeds(id, created_at, updated_at, feed_name, url_name, user_id)
values($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, feed_name, url_name, user_id
`

type CreateFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedName  string
	UrlName   string
	UserID    uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedName,
		arg.UrlName,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedName,
		&i.UrlName,
		&i.UserID,
	)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
select id, created_at, updated_at, feed_name, url_name, user_id from feeds
`

func (q *Queries) GetFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedName,
			&i.UrlName,
			&i.UserID,
		); err != nil {
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