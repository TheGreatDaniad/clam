package graphql

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"podcasts": &graphql.Field{
				Type: graphql.NewList(podcastType),
				Args: graphql.FieldConfigArgument{
					"search": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"categoryName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: ResolvePodcasts,
			},
		},
	},
)

var GraphqlHandler = func(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: QueryType,
	})
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  params.Query,
		VariableValues: params.Variables,
		OperationName:  params.OperationName,
	})

	json.NewEncoder(w).Encode(result)
}
