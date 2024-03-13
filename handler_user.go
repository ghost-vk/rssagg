package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ghost-vk/rssagg/internal/auth"
	"github.com/ghost-vk/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	dbUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(dbUser))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, "forbidden")
		return
	}

	// Context() позволяет следить за стейтом нескольких goroutines
	// Важная функция - это отмена действий, например, дает возможность
	// закончить http request
	dbUser, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Couldn't get user info: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(dbUser))
}
