package main

import (
	"fmt"
	"net/http"
	"server/internal/auth"
	"server/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (cfg *apiConfig) middleWareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Could not get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
