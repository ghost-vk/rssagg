package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ghost-vk/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerGetUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Offset int32 `json:"offset"`
		Limit  int32 `json:"limit"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		ApiKey: user.ApiKey,
		Offset: params.Offset,
		Limit:  params.Limit,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondWithJSON(w, 200, databasePostsToPost(posts))
}
