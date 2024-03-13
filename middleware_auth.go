package main

import (
	"fmt"
	"net/http"

	"github.com/ghost-vk/rssagg/internal/auth"
	"github.com/ghost-vk/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		handler(w, r, dbUser)
	}
}
