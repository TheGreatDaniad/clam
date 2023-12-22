package services

// here I implement a very basic service registry system

const (
	REST_CONNECTION_MODE = iota
	GRPC_CONNECTION_MODE
)
const (
	PODCAST_SERVICE = "PODCAST_SERVICE"
)

// this is a struct representing the data of each microservice
// used to track the important data of each microservice
type ServiceData struct {
	Name           string
	URL            string
	ConnectionMode int
	Version        string
	HealthCheckURL string
	Metadata       map[string]string
}

type Services struct {
	PodcastService *PodcastService
}

func GetServices() *Services {
	return &Services{
		PodcastService: &PodcastService{
			ServiceData: ServiceData{
				Name:           "Podcast Service",
				URL:            "https://601f1754b5a0e9001706a292.mockapi.io/podcasts",
				ConnectionMode: REST_CONNECTION_MODE,
				Version:        "1.0.0",
			},
		},
	}
}