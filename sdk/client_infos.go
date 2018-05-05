package sdk

import (
	"github.com/dghubble/sling"
	"github.com/gpaulo00/gh0st/models"
)

// Infos makes a fluent api to manage infos
type Infos struct {
	client *sling.Sling
}

// Infos returns a Infos, to manage infos
func (c *Client) Infos() *Infos {
	return &Infos{client: c.client}
}

// List returns a list of infos in the server
func (c *Infos) List() (res []*models.Info, err error) {
	var bad models.ErrorResult
	_, err = c.client.Get(models.InfoPath.String()).Receive(&res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return
}

// Add inserts a new workspace into the server
func (c *Infos) Add(obj *models.Info) (*models.Info, error) {
	var bad models.ErrorResult
	res := new(models.Info)
	_, err := c.client.Post(models.InfoPath.String()).
		BodyJSON(obj).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}

// Delete removes a workspace from the server
func (c *Infos) Delete(id uint64) error {
	var bad models.ErrorResult
	res := new(models.DoneResult)
	_, err := c.client.Delete(models.InfoPath.WithID(id)).Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return err
}

// Get returns a workspace from the server
func (c *Infos) Get(id uint64) (*models.Info, error) {
	var bad models.ErrorResult
	res := new(models.Info)
	_, err := c.client.Get(models.InfoPath.WithID(id)).Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}
