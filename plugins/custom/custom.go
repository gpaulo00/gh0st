package custom

import (
	"fmt"

	"github.com/gpaulo00/gh0st/models"
)

// ImportSource contains information about an imported source
type ImportSource struct {
	Source models.Source `json:"source" binding:"required"`
	Hosts  []ImportHost  `json:"hosts" binding:"required"`
}

// ImportHost contains information about an imported host
type ImportHost struct {
	Host     models.Host      `json:"host" binding:"required"`
	Notes    []models.Note    `json:"notes"`
	Issues   []models.Issue   `json:"issues"`
	Services []models.Service `json:"services"`
}

// ImportResult is the result of an import
type ImportResult struct {
	Hosts    int `json:"hosts"`
	Services int `json:"services"`
	Notes    int `json:"notes"`
	Issues   int `json:"issues"`
}

func (i ImportResult) String() string {
	return fmt.Sprintf(
		"imported data: hosts = %d, services = %d, notes = %d, issues = %d",
		i.Hosts, i.Services, i.Notes, i.Issues,
	)
}
