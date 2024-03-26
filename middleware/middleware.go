package middleware

import "net/http"

type Middlware func(http.Handler) http.Handler

func CombineMiddlewares(middleware ...Middlware) Middlware {
	return func(next http.Handler) http.Handler {
		for i := range middleware {
			x := middleware[i]
			next = x(next)
		}
		return next
	}
}
