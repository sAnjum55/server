-- name: CreateFeed :one
insert into feeds(id, created_at, updated_at, feed_name, url_name, user_id)
values($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
select * from feeds;