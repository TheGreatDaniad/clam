package main

import (
	"net/http"
	"golang.org/x/time/rate"
	"github.com/friendsofgo/graphiql"
	gql "github.com/thegreatdaniad/clam/graphql"
	"github.com/rs/cors"
)

func limit(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(1, 3) 

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, 
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		Debug:          true, 
	})

	http.Handle("/graphql", limit(c.Handler(http.HandlerFunc(gql.GraphqlHandler))))
	http.Handle("/graphiql", graphiqlHandler)

	http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8080", nil)
}
