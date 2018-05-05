package sdk

import (
	"github.com/dghubble/sling"
	"github.com/gpaulo00/gh0st/models"
)

// Sources makes a fluent api to manage sources
type Sources struct {
	client *sling.Sling
}

// Sources returns a Sources, to manage sources
func (c *Client) Sources() *Sources {
	return &Sources{client: c.client}
}

// List returns a list of sources in the server
func (c *Sources) List() (res []*models.Source, err error) {
	var bad models.ErrorResult
	_, err = c.client.Get(models.SourcePath.String()).Receive(&res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return
}

// Add inserts a new workspace into the server
func (c *Sources) Add(obj *models.Source) (*models.Source, error) {
	var bad models.ErrorResult
	res := new(models.Source)
	_, err := c.client.Post(models.SourcePath.String()).
		BodyJSON(obj).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}

// Delete removes a workspace from the server
func (c *Sources) Delete(id uint64) error {
	var bad models.ErrorResult
	res := new(models.DoneResult)
	_, err := c.client.Delete(models.SourcePath.WithID(id)).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return err
}

// Get returns a workspace from the server
func (c *Sources) Get(id uint64) (*models.Source, error) {
	var bad models.ErrorResult
	res := new(models.Source)
	_, err := c.client.Get(models.SourcePath.WithID(id)).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}
