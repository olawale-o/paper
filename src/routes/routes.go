package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {

	// Target service URL
	targetURL, err := url.Parse("http://localhost:8083")     // Replace with your service URL
	authTargetURL, err := url.Parse("http://localhost:8081") // Replace with your service URL
	if err != nil {
		panic(err)
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	authProxy := httputil.NewSingleHostReverseProxy(authTargetURL)

	r.Route("/api/v1", func(r chi.Router) {

		r.Handle("/articles/*", http.StripPrefix("/api/v1/articles", proxy)) // Route /api/* to the service
		r.Handle("/auth/*", http.StripPrefix("/api/v1/auth", authProxy))     // Route /api/* to the service
	})

}
