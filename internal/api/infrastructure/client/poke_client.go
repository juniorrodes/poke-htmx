package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

type PokeClient struct {
	client http.Client
}

const (
	pokeApiPath = "https://pokeapi.co/api"
)

func NewPokeClient() *PokeClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return &PokeClient{
		client: http.Client{
			Transport: transport,
		},
	}
}

func (c *PokeClient) DoRequest(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", pokeApiPath, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
