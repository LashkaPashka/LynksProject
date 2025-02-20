package middleware

import (
	"ShorteNer/configs"
	"ShorteNer/pkg/jwt"
	"context"
	"net/http"
	"strings"
)
type myString string
const (
	emailKey myString = "emailKey"
)

func WriteStatusCode(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, conf *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			WriteStatusCode(w)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		
		isValid, data := jwt.NewJWT(conf.Secret).Parse(token)
		if !isValid {
			WriteStatusCode(w)
		}

		var ctx = context.WithValue(r.Context(), emailKey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}