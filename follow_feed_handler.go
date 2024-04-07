package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/sAnjum55/server/internal/database"
)

func (apiCfg *apiConfig) createFollowFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	follow_feeds, err := apiCfg.DB.CreateFollowFeeds(r.Context(), database.CreateFollowFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("could not follow feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFollowFeedToFollowFeed(follow_feeds))
}

func (apiCfg *apiConfig) getFollowFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollowed, err := apiCfg.DB.GetFeedsForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 404, "User does not follow any feeds")
		return
	}
	respondWithJSON(w, 200, databaseFollowFeedsToFollowFeeds(feedsFollowed))
}

func (apiCfg *apiConfig) deleteFollowFeedForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDString := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDString)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not parse the uql param: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFollowFeedForUser(r.Context(), database.DeleteFollowFeedForUserParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not delete: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})

}
