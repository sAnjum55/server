-- +goose Up

create table feeds(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    feed_name text not null,
    url_name text unique not null,
    user_id uuid not null references users(id) on delete cascade
);

-- +goose Down
drop table feeds;