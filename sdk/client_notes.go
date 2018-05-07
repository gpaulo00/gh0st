package sdk

import (
	"github.com/dghubble/sling"
	"github.com/gpaulo00/gh0st/models"
)

// Notes makes a fluent api to manage notes
type Notes struct {
	client *sling.Sling
}

// Notes returns a Notes, to manage notes
func (c *Client) Notes() *Notes {
	return &Notes{client: c.client}
}

// List returns a list of notes in the server
func (c *Notes) List() (res []*models.Note, err error) {
	var bad models.ErrorResult
	_, err = c.client.Get(models.NotePath.String()).Receive(&res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return
}

// Add inserts a new workspace into the server
func (c *Notes) Add(obj *models.Note) (*models.Note, error) {
	var bad models.ErrorResult
	res := new(models.Note)
	_, err := c.client.Post(models.NotePath.String()).
		BodyJSON(obj).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}

// Delete removes a workspace from the server
func (c *Notes) Delete(id uint64) error {
	var bad models.ErrorResult
	res := new(models.DoneResult)
	_, err := c.client.Delete(models.NotePath.WithID(id)).Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return err
}

// Get returns a workspace from the server
func (c *Notes) Get(id uint64) (*models.Note, error) {
	var bad models.ErrorResult
	res := new(models.Note)
	_, err := c.client.Get(models.NotePath.WithID(id)).Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}
