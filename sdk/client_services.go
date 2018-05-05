package sdk

import (
	"github.com/dghubble/sling"
	"github.com/gpaulo00/gh0st/models"
)

// Services makes a fluent api to manage services
type Services struct {
	client *sling.Sling
}

// Services returns a Services, to manage services
func (c *Client) Services() *Services {
	return &Services{client: c.client}
}

// List returns a list of services in the server
func (c *Services) List() (res []*models.Host, err error) {
	var bad models.ErrorResult
	_, err = c.client.Get(models.ServicePath.String()).Receive(&res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return
}

// Add inserts a new workspace into the server
func (c *Services) Add(obj *models.Host) (*models.Host, error) {
	var bad models.ErrorResult
	res := new(models.Host)
	_, err := c.client.Post(models.ServicePath.String()).
		BodyJSON(obj).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}

// Delete removes a workspace from the server
func (c *Services) Delete(id uint64) error {
	var bad models.ErrorResult
	res := new(models.DoneResult)
	_, err := c.client.Delete(models.ServicePath.WithID(id)).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return err
}

// Get returns a workspace from the server
func (c *Services) Get(id uint64) (*models.Host, error) {
	var bad models.ErrorResult
	res := new(models.Host)
	_, err := c.client.Get(models.ServicePath.WithID(id)).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}
