package pokeapi

import (
	"net/http"
	"time"
)

type CliConfig struct {
	CachedInfo Cache
	Pokedex    Pokedex
	Next       string
	Previous   string
}

type Client struct {
	httpClient http.Client
	Config     CliConfig
}

func NewClient() Client {
	return Client{
		httpClient: *http.DefaultClient,
		Config: CliConfig{
			Pokedex:    NewPokedex(),
			CachedInfo: NewCache(5 * time.Second),
		},
	}
}
