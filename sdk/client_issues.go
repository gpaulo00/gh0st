package sdk

import (
	"github.com/dghubble/sling"
	"github.com/gpaulo00/gh0st/models"
)

// Issues makes a fluent api to manage issues
type Issues struct {
	client *sling.Sling
}

// Issues returns a Issues, to manage issues
func (c *Client) Issues() *Issues {
	return &Issues{client: c.client}
}

// List returns a list of issues in the server
func (c *Issues) List() (res []*models.Issue, err error) {
	var bad models.ErrorResult
	_, err = c.client.Get(models.IssuePath.String()).Receive(&res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return
}

// Add inserts a new workspace into the server
func (c *Issues) Add(obj *models.Issue) (*models.Issue, error) {
	var bad models.ErrorResult
	res := new(models.Issue)
	_, err := c.client.Post(models.IssuePath.String()).
		BodyJSON(obj).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}

// Delete removes a workspace from the server
func (c *Issues) Delete(id uint64) error {
	var bad models.ErrorResult
	res := new(models.DoneResult)
	_, err := c.client.Delete(models.IssuePath.WithID(id)).Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return err
}

// Get returns a workspace from the server
func (c *Issues) Get(id uint64) (*models.Issue, error) {
	var bad models.ErrorResult
	res := new(models.Issue)
	_, err := c.client.Get(models.IssuePath.WithID(id)).Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}
