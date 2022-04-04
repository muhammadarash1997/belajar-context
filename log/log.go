package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

const requestIDKey = 42

func Println(ctx context.Context, msg string) {
	// We need to get a request ID key
	// ctx.Value should return id or nil
	id, ok := ctx.Value(ctx).(int64)
	if !ok {
		log.Println("could not find request ID in context")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

// Wrapper
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Basically what we are doing in here is receiving a context then adding a value then sending it back

		// Get context from the request
		ctx := r.Context()
		// Generate a new number with random function and that's going to be our id
		id := rand.Int63()
		// Put id in the context
		ctx = context.WithValue(ctx, requestIDKey, id)
		// Send it back
		f(w, r.WithContext(ctx))
	}
}
