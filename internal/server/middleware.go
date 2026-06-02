package server

import "net/http"

// Middleware represents our standard adapter signature
type Middleware func(http.Handler) http.Handler

// Chain constructs our "Russian Nesting Doll" handler stack at server startup.
func Chain(core http.Handler, middlewares ...Middleware) http.Handler {
	// Loop backwards to wrap them from inside-out so they execute outside-in
	for i := len(middlewares) - 1; i >= 0; i-- {
		core = middlewares[i](core)
	}
	return core
}
