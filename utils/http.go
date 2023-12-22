package utils

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/sony/gobreaker"
	"github.com/spf13/viper"
)

func Request(ctx context.Context, cb *gobreaker.CircuitBreaker, method, url string, body io.Reader) ([]byte, error) {
	var responseBody []byte
	var err error
	// we are using a circuit breaker to avoid cascading failures in each microservice
	// if a microservice is down, we will not try to connect to it again until a certain amount of time has passed
	_, err = cb.Execute(func() (interface{}, error) {
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			return nil, err
		}
		req = req.WithContext(ctx)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		responseBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	return responseBody, err
}

func StartLogger() {
	filePath := viper.GetString("logFilePath")
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file")
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}
