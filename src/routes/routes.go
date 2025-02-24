package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {

	// Target service URL
	authTargetURL, err := url.Parse("http://localhost:8081")    // Replace with your service URL
	authorTargetURL, err := url.Parse("http://localhost:8082")  // Replace with your service URL
	articleTargetURL, err := url.Parse("http://localhost:8083") // Replace with your service URL
	commentTargetURL, err := url.Parse("http://localhost:8084") // Replace with your service URL
	if err != nil {
		panic(err)
	}

	// Create a reverse proxy
	articleProxy := httputil.NewSingleHostReverseProxy(articleTargetURL)
	authProxy := httputil.NewSingleHostReverseProxy(authTargetURL)
	authorProxy := httputil.NewSingleHostReverseProxy(authorTargetURL)
	commentProxy := httputil.NewSingleHostReverseProxy(commentTargetURL)

	r.Route("/api/v1", func(r chi.Router) {

		r.Handle("/articles/*", http.StripPrefix("/api/v1/articles", articleProxy)) // Route /api/* to the service
		r.Handle("/auth/*", http.StripPrefix("/api/v1/auth", authProxy))            // Route /api/* to the service     // Route /api/* to the service
		r.Handle("/authors/*", http.StripPrefix("/api/v1/authors", authorProxy))    // Route /api/* to the service
		r.Handle("/comments/*", http.StripPrefix("/api/v1/comments", commentProxy)) // Route /api/* to the service

	})

}
