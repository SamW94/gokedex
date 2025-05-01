package pokeapi

import (
	"net/http"
	"time"

	"github.com/SamW94/gokedex/internal/pokecache"
)

type WebClient struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) WebClient {
	return WebClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
