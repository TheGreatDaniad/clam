package graphql

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/graphql-go/graphql"
	"github.com/thegreatdaniad/clam/services"
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

func ResolvePodcasts(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context
	search, _ := p.Args["search"].(string)
	title, _ := p.Args["title"].(string)
	categoryName, _ := p.Args["categoryName"].(string)
	page, _ := p.Args["page"].(int)
	limit, _ := p.Args["limit"].(int)

	// implementing some basic input validations 
	maxLength := "200"

	if !govalidator.StringLength(search, "", maxLength) {
		return nil, errors.New("search string exceeds maximum length of 200 characters")
	}
	if !govalidator.StringLength(title, "", maxLength) {
		return nil, errors.New("title string exceeds maximum length of 200 characters")
	}
	if !govalidator.StringLength(categoryName, "", maxLength) {
		return nil, errors.New("category name string exceeds maximum length of 200 characters")
	}
	qp := services.QueryPodcast{
		Search:       search,
		Title:        title,
		CategoryName: categoryName,
		Page:         page,
		Limit:        limit,
	}
	s := services.GetServices()
	podcasts, err := s.PodcastService.GetPodcasts(ctx, qp)
	if err != nil {
		return nil, err
	}

	return podcasts, nil
}
