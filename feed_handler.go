package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sAnjum55/server/internal/database"
)

func (apiCfg *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedName string `json:"feed_name"`
		UrlName  string `json:"url_name"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error reading params: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedName:  params.FeedName,
		UrlName:   params.UrlName,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("could not create feed: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("No feeds found: %v", err))
		return
	}
	var formattedFeeds = []Feed{}
	for _, val := range feeds {
		formattedFeeds = append(formattedFeeds, databaseFeedToFeed(val))
	}
	respondWithJSON(w, 200, formattedFeeds)
}
