package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/meziaris/gofinance/internal/pkg/handler"
	"github.com/meziaris/gofinance/internal/pkg/reason"
)

type AccessTokenVerifier interface {
	VerifyAccessToken(tokenString string) (sub string, err error)
}

func AuthMiddleware(tokenCreator AccessTokenVerifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken := tokenFromHeader(r)
			if accessToken == "" {
				handler.ResponseError(w, http.StatusUnprocessableEntity, reason.Unauthorized)
				return
			}
			sub, err := tokenCreator.VerifyAccessToken(accessToken)
			if err != nil {
				handler.ResponseError(w, http.StatusUnprocessableEntity, reason.Unauthorized)
				return
			}

			// create new context from `r` request context, and assign key `"user_id"`
			// to value of `sub"`
			ctx := context.WithValue(r.Context(), "user_id", sub)

			// call the next handler in the chain, passing the response writer and
			// the updated request object with the new context value.
			//
			// note: context.Context values are nested, so any previously set
			// values will be accessible as well, and the new `"user_id"` key
			// will be accessible from this point forward.
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func tokenFromHeader(r *http.Request) string {
	var accessToken string
	bearerToken := r.Header.Get("Authorization")
	field := strings.Fields(bearerToken)

	if len(field) != 0 && field[0] == "Bearer" {
		accessToken = field[1]
	}

	return accessToken
}
