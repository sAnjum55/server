package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/sAnjum55/server/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedName  string    `json:"feed_name"`
	UrlName   string    `json:"url_name"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		FeedName:  dbFeed.FeedName,
		UrlName:   dbFeed.UrlName,
		UserID:    dbFeed.UserID,
	}

}

type FollowFeed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFollowFeedToFollowFeed(dbFollowFeeds database.FollowFeed) FollowFeed {
	return FollowFeed{
		ID:        dbFollowFeeds.ID,
		CreatedAt: dbFollowFeeds.CreatedAt,
		UpdatedAt: dbFollowFeeds.UpdatedAt,
		UserID:    dbFollowFeeds.UserID,
		FeedID:    dbFollowFeeds.FeedID,
	}
}

func databaseFollowFeedsToFollowFeeds(dbFollowFeeds []database.FollowFeed) []FollowFeed {
	var formattedFollowFeeds = []FollowFeed{}
	for _, val := range dbFollowFeeds {
		formattedFollowFeeds = append(formattedFollowFeeds, databaseFollowFeedToFollowFeed(val))
	}
	return formattedFollowFeeds
}
