package models

// IssuesStats is information about issues,
// that is used to create charts.
type IssuesStats struct {
	ID       int8   `json:"-"`
	Critical uint64 `json:"critical"`
	High     uint64 `json:"high"`
	Medium   uint64 `json:"medium"`
	Low      uint64 `json:"low"`
	Info     uint64 `json:"info"`
}
