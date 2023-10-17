package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// reqKey is a type used for context keys.
type reqKey int

// RequestIDKey is a key for 'request id' context value.
const RequestIDKey reqKey = 123

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := uuid.NewString() // creating the req id

		ctx := context.WithValue(r.Context(), RequestIDKey, token) // creating a new request contest

		reqID, ok := ctx.Value(RequestIDKey).(string)
		if !ok {
			reqID = "Unknown"
		}

		// Ensure log of completion when the request is done.
		defer log.Println("completed", reqID)

		// Pass the responsibility to the next HandlerFunc in the chain.
		next(w, r)

	}
}
