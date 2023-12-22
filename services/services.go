package services

import "log"

const (
	REST_CONNECTION_MODE = iota
	GRPC_CONNECTION_MODE
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

// this is an interface representing each of microservices
// all of the microservices must implement this common functions
// they generally provice general control and access to the microservices (like obtaining some metadata)
type IService interface {
	// check connection health to the microservice
	CheckConnection() (bool, error)
	GetServiceData() *ServiceData
}

type Services []IService

func CheckServicesHealth(services Services) (bool, error) {
	ok := true
	for _, service := range services {
		working, err := service.CheckConnection()
		log.Println("Checking service health: ", service.GetServiceData().Name)

		if err != nil {
			return false, err
		}
		if !working {
			ok = false
		}
	}
	return ok, nil
}
