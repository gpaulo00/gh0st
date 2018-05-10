package models

// IssuesSummary is information about issues,
// that is used to create charts.
type IssuesSummary struct {
	ID    int8 `json:"-"`
	Stats struct {
		Critical uint64 `json:"critical"`
		High     uint64 `json:"high"`
		Medium   uint64 `json:"medium"`
		Low      uint64 `json:"low"`
		Info     uint64 `json:"info"`
	} `json:"stats"`
	Titles []struct {
		Title  string     `json:"title"`
		Level  IssueLevel `json:"level"`
		Number uint64     `json:"number"`
	} `json:"summary"`
}
