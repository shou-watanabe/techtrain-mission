package middleware

import (
	"net/http"

	"go-gacha-server/src/core/context"
)

func Context(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewContext(r)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
