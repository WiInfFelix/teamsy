package auth

import (
	"context"
	"net/http"
	"teamsy/internal/pkg/jwt"
	"teamsy/internal/pkg/users"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func AuthMiddleware() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			header := request.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(writer, request)
				return
			}

			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(writer, "Invalid authorization token", http.StatusForbidden)
				return
			}

			user := users.User{Username: username}
			id, err := users.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(writer, request)
				return
			}

			user.ID = uint(id)

			ctx := context.WithValue(request.Context(), userCtxKey, &user)

			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request)
		})
	}
}

func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
