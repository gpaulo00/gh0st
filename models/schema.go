package models

import (
	"net"
	"time"
)

// JSON represents schema-less data in the database
type JSON map[string]interface{}

// Workspace organizes all the data
type Workspace struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `sql:"default:now()" json:"createdAt"`
}

// Source is the source of data
type Source struct {
	ID          uint64     `json:"id,omitempty"`
	WorkspaceID uint64     `sql:"workspace_id" json:"workspace" binding:"required"`
	Generator   string     `json:"generator" binding:"required"`
	SourceInfo  *JSON      `sql:"source_info" json:"sourceInfo"`
	GeneratedAt *time.Time `sql:"generated_at" json:"generatedAt"`
	CreatedAt   time.Time  `sql:"default:now()" json:"createdAt,omitempty"`
}

// Host is a target of data
type Host struct {
	ID        uint64    `json:"id,omitempty"`
	SourceID  uint64    `sql:"source_id" json:"source"`
	Address   net.IP    `json:"address" binding:"required"`
	State     string    `json:"state" binding:"required"`
	CreatedAt time.Time `sql:"default:now()" json:"createdAt,omitempty"`
}

// Service is running service of a host
type Service struct {
	ID        uint64    `json:"id,omitempty"`
	HostID    uint64    `sql:"host_id" json:"host"`
	Protocol  string    `json:"protocol" binding:"required"`
	Port      uint16    `json:"port" binding:"required"`
	State     string    `json:"state" binding:"required"`
	Service   *string   `json:"service"`
	CreatedAt time.Time `sql:"default:now()" json:"createdAt,omitempty"`
}

// Note is an extra information of a host
type Note struct {
	ID        uint64    `json:"id,omitempty"`
	HostID    uint64    `sql:"host_id" json:"host"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `sql:"default:now()" json:"createdAt,omitempty"`
}

// Issue represents an issue in a host
type Issue struct {
	ID        uint64     `json:"id,omitempty"`
	HostID    uint64     `sql:"host_id" json:"host"`
	Level     IssueLevel `json:"level" binding:"required"`
	Title     string     `json:"title" binding:"required"`
	Content   string     `json:"content" binding:"required"`
	CreatedAt time.Time  `sql:"default:now()" json:"createdAt,omitempty"`
}

// IssueLevel represents the level of an issue
type IssueLevel string

const (
	// CriticalIssue is the "critical" level of an issue
	CriticalIssue = IssueLevel("critical")

	// HighIssue is the "high" level of an issue
	HighIssue = IssueLevel("high")

	// MediumIssue is the "medium" level of an issue
	MediumIssue = IssueLevel("medium")

	// LowIssue is the "low" level of an issue
	LowIssue = IssueLevel("low")

	// InfoIssue is the "info" level of an issue
	InfoIssue = IssueLevel("info")
)
