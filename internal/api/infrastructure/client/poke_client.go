package client

import (
	"crypto/tls"
	"net/http"

	"github.com/juniorrodes/poke-htmx/internal/api/app/entities"
)

type PokeClient[T entities.Entity] struct {
	client http.Client
}

func NewPokeClient[T entities.Entity]() *PokeClient[T] {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return &PokeClient[T]{
		client: http.Client{
			Transport: transport,
		},
	}
}

func (c *PokeClient[T]) DoRequest(endpoint string) T {

}
