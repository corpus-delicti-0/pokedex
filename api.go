package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetchData(cfg *config, url string) ([]byte, error) {
	data, found := cfg.cache.Get(url)
	if found {
		return data, nil
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", response.Status)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	cfg.cache.Add(url, data)

	return data, nil
}
