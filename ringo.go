// Package ringo provides an unofficial client to the unpublished Ring API.
package ringo

import (
	"net/http"
	"net/url"
)

// Ringo is used to make requests to the unpublished Ring API.
type Ringo struct {
	BaseURL *url.URL
	Client  *http.Client
	Config  *Config
}

// NewRingo constructs a new Ringo which can make requests to the unpublished Ring API.
func NewRingo(baseURL *url.URL, client *http.Client, config *Config) *Ringo {
	if client == nil {
		client = http.DefaultClient
	}

	return &Ringo{
		BaseURL: baseURL,
		Client:  client,
		Config:  config,
	}
}
