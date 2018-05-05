package sdk

import "github.com/gpaulo00/gh0st/models"

// Import imports a new source with data into the database
func (c *Client) Import(data *models.ImportForm) (*models.ImportResult, error) {
	var bad models.ErrorResult
	res := new(models.ImportResult)
	_, err := c.client.Post(models.ImportPath).
		BodyJSON(data).
		Receive(res, &bad)

	if err == nil {
		err = bad.Err()
	}
	return res, err
}
