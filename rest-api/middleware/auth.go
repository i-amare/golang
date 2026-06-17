package middleware

import (
	"context"
	"net/http"

	"github.com/i-amare/rest-api/utils"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			utils.WriteError(w, http.StatusUnauthorized, "No auth token provided", nil)
			return
		}

		userID, err := utils.VerifyAuthToken(token)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, "Not authorised", err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "UserID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(r *http.Request) int64 {
	value := r.Context().Value("UserID")
	if userID, ok := value.(int64); ok {
		return userID
	}
	return -1
}
