package sdk

import (
	"github.com/dghubble/sling"
	"github.com/gpaulo00/gh0st/models"
)

// Workspaces makes a fluent api to manage workspaces
type Workspaces struct {
	client *sling.Sling
}

// Workspaces returns a Workspaces, to manage workspaces
func (c *Client) Workspaces() *Workspaces {
	return &Workspaces{client: c.client}
}

// List returns a list of workspaces in the server
func (c *Workspaces) List() (res []*models.Workspace, err error) {
	var bad models.ErrorResult
	_, err = c.client.Get(models.WorkspacePath.String()).Receive(&res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return
}

// Add inserts a new workspace into the server
func (c *Workspaces) Add(obj *models.Workspace) (*models.Workspace, error) {
	var bad models.ErrorResult
	res := new(models.Workspace)
	_, err := c.client.Post(models.WorkspacePath.String()).
		BodyJSON(obj).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}

// Delete removes a workspace from the server
func (c *Workspaces) Delete(id uint64) error {
	var bad models.ErrorResult
	res := new(models.DoneResult)
	_, err := c.client.Path(models.WorkspacePath.WithID(id)).Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return err
}

// Get returns a workspace from the server
func (c *Workspaces) Get(id uint64) (*models.Workspace, error) {
	var bad models.ErrorResult
	res := new(models.Workspace)
	_, err := c.client.Get(models.WorkspacePath.WithID(id)).Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}

// Update modifies a workspace from the server
func (c *Workspaces) Update(id uint64, obj *models.Workspace) (*models.Workspace, error) {
	var bad models.ErrorResult
	res := new(models.Workspace)
	_, err := c.client.Put(models.WorkspacePath.WithID(id)).
		BodyJSON(obj).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}
