package pokeapi

import (
	"net/http"
	"time"

	"github.com/carloscfgos1980/pokodex/internal/pokecache"
)

// Client represents a client for interacting with the PokeAPI. It contains an HTTP client with a configurable timeout.
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// NewClient creates a new PokeAPI client with the specified timeout for HTTP requests.
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
