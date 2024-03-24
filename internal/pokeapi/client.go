package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

// make an active client for given timeout duration
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
