package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/sony/gobreaker"
)

var imagesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Images",
		Fields: graphql.Fields{
			"default": &graphql.Field{
				Type: graphql.String,
			},
			"featured": &graphql.Field{
				Type: graphql.String,
			},
			"thumbnail": &graphql.Field{
				Type: graphql.String,
			},
			"wide": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var podcastType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Podcast",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"images": &graphql.Field{
				Type: imagesType,
			},
			"isExclusive": &graphql.Field{
				Type: graphql.Boolean,
			},
			"publisherName": &graphql.Field{
				Type: graphql.String,
			},
			"publisherId": &graphql.Field{
				Type: graphql.String,
			},
			"mediaType": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"categoryId": &graphql.Field{
				Type: graphql.String,
			},
			"categoryName": &graphql.Field{
				Type: graphql.String,
			},
			"hasFreeEpisodes": &graphql.Field{
				Type: graphql.Boolean,
			},
			"playSequence": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func resolvePodcasts(p graphql.ResolveParams) (interface{}, error) {
	// Extract query parameters
	search, _ := p.Args["search"].(string)
	title, _ := p.Args["title"].(string)
	categoryName, _ := p.Args["categoryName"].(string)
	page, _ := p.Args["page"].(int)
	limit, _ := p.Args["limit"].(int)

	// Construct the query parameters
	queryParams := url.Values{}
	if search != "" {
		queryParams.Add("search", search)
	}
	if title != "" {
		queryParams.Add("title", title)
	}
	if categoryName != "" {
		queryParams.Add("categoryName", categoryName)
	}
	if page != 0 {
		queryParams.Add("page", strconv.Itoa(page))
	}
	if limit != 0 {
		queryParams.Add("limit", strconv.Itoa(limit))
	}



	return podcasts, nil
}
