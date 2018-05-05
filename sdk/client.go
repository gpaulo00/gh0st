package sdk

import (
	"net/url"

	"github.com/dghubble/sling"
)

// Client holds information to connect to a gh0st server
type Client struct {
	client *sling.Sling
}

// New builds a new client connection to the RESTful API
func New(api string) (*Client, error) {
	if _, err := url.Parse(api); err != nil {
		return nil, err
	}

	c := &Client{
		client: sling.New().Base(api),
	}
	return c, nil
}
