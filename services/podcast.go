package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sony/gobreaker"
	"github.com/thegreatdaniad/clam/utils"
)

type Podcast struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Images struct {
		Default   string `json:"default"`
		Featured  string `json:"featured"`
		Thumbnail string `json:"thumbnail"`
		Wide      string `json:"wide"`
	} `json:"images"`
	IsExclusive     bool   `json:"isExclusive"`
	PublisherName   string `json:"publisherName"`
	PublisherID     string `json:"publisherId"`
	MediaType       string `json:"mediaType"`
	Description     string `json:"description"`
	CategoryID      string `json:"categoryId"`
	CategoryName    string `json:"categoryName"`
	HasFreeEpisodes bool   `json:"hasFreeEpisodes"`
	PlaySequence    string `json:"playSequence"`
}

type PodcastService struct {
	ServiceData
}

func (p *PodcastService) CheckConnection() (bool, error) {
	return true, nil
}
func (p *PodcastService) GetServiceData() (*ServiceData, error) {
	return &p.ServiceData, nil
}

func (p *PodcastService) GetPodcast(ctx context.Context, id string) (*Podcast, error) {
	return nil, nil
}

func (p *PodcastService) GetPodcasts(ctx context.Context, queryParams url.Values) ([]*Podcast, error) {
	// Construct the URL with query parameters
	requestURL := fmt.Sprintf("https://601f1754b5a0e9001706a292.mockapi.io/podcasts?%s", queryParams.Encode())

	// Use the Request function to make the API call
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{}) // Assuming a circuit breaker is set up
	responseBytes, err := Request(context.Background(), cb, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	// Parse the response
	var podcasts []Podcast
	err = json.Unmarshal(responseBytes, &podcasts)
	if err != nil {
		return nil, err
	}
}
