package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

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

type QueryPodcast struct {
	Search       string `json:"search"`
	Title        string `json:"title"`
	CategoryName string `json:"categoryName"`
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
}
type PodcastService struct {
	ServiceData
}


func (p *PodcastService) CheckConnection() (bool, error) {
	return true, nil
}
func (p *PodcastService) GetServiceData() *ServiceData {
	return &p.ServiceData
}

func (p *PodcastService) GetPodcast(ctx context.Context, id string) (*Podcast, error) {
	return nil, nil
}

func (p *PodcastService) GetPodcasts(ctx context.Context, qp QueryPodcast) ([]Podcast, error) {
	if p.ConnectionMode == REST_CONNECTION_MODE {
		queryParams := url.Values{}
		if qp.Search != "" {
			queryParams.Add("search", qp.Search)
		}
		if qp.Title != "" {
			queryParams.Add("title", qp.Title)
		}
		if qp.CategoryName != "" {
			queryParams.Add("categoryName", qp.CategoryName)
		}
		if qp.Page != 0 {
			queryParams.Add("page", strconv.Itoa(qp.Page))
		}
		if qp.Limit != 0 {
			queryParams.Add("limit", strconv.Itoa(qp.Limit))
		}

		requestURL := fmt.Sprintf("https://601f1754b5a0e9001706a292.mockapi.io/podcasts?%s", queryParams.Encode())

		cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{})
		responseBytes, err := utils.Request(context.Background(), cb, "GET", requestURL, nil)
		if err != nil {
			return nil, err
		}
		var podcasts []Podcast
		err = json.Unmarshal(responseBytes, &podcasts)
		if err != nil {
			return nil, err
		}
		return podcasts, nil
	}
	if p.ConnectionMode == GRPC_CONNECTION_MODE {
		// not implemented yet
		return nil, nil
	}

	return nil, nil
}
