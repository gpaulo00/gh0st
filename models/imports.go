package models

import "fmt"

// ImportForm contains information about an imported source
type ImportForm struct {
	Source Source `json:"source" binding:"required"`
	Hosts  []struct {
		Host     Host      `json:"host" binding:"required"`
		Infos    []Info    `json:"infos"`
		Services []Service `json:"services"`
	} `json:"hosts" binding:"required"`
}

// ImportResult is the result of an import
type ImportResult struct {
	Hosts    int `json:"hosts"`
	Services int `json:"services"`
	Infos    int `json:"infos"`
}

func (i ImportResult) String() string {
	return fmt.Sprintf(
		"imported data: hosts = %d, services = %d, infos = %d",
		i.Hosts, i.Services, i.Infos,
	)
}
