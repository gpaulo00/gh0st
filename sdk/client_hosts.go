package sdk

import (
	"github.com/dghubble/sling"
	"github.com/gpaulo00/gh0st/models"
)

// Hosts makes a fluent api to manage hosts
type Hosts struct {
	client *sling.Sling
}

// Hosts returns a Hosts, to manage hosts
func (c *Client) Hosts() *Hosts {
	return &Hosts{client: c.client}
}

// List returns a list of hosts in the server
func (c *Hosts) List() (res []*models.Host, err error) {
	var bad models.ErrorResult
	_, err = c.client.Get(models.HostPath.String()).Receive(&res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return
}

// Add inserts a new workspace into the server
func (c *Hosts) Add(obj *models.Host) (*models.Host, error) {
	var bad models.ErrorResult
	res := new(models.Host)
	_, err := c.client.Post(models.HostPath.String()).
		BodyJSON(obj).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}

// Delete removes a workspace from the server
func (c *Hosts) Delete(id uint64) error {
	var bad models.ErrorResult
	res := new(models.DoneResult)
	_, err := c.client.Delete(models.HostPath.WithID(id)).Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return err
}

// Get returns a workspace from the server
func (c *Hosts) Get(id uint64) (*models.Host, error) {
	var bad models.ErrorResult
	res := new(models.Host)
	_, err := c.client.Get(models.HostPath.WithID(id)).Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}
