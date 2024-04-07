-- name: CreateFollowFeeds :one

insert into follow_feed(id, created_at, updated_at, user_id,feed_id)
values($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedsForUser :many

select *
from follow_feed
where user_id = $1;

-- name: DeleteFollowFeedForUser :exec
delete 
from follow_feed
where id=$1 AND user_id=$2;