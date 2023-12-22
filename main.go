package main

import (
	"net/http"

	"github.com/friendsofgo/graphiql"
	gql "github.com/thegreatdaniad/clam/graphql"
)

func main() {

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/graphql", gql.GraphqlHandler)
	http.Handle("/graphiql", graphiqlHandler)
	http.ListenAndServe(":8080", nil)

}
