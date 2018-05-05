package models

import (
	"net"
	"time"

	"github.com/go-pg/pg/orm"
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

// Info is an extra information of a host
type Info struct {
	ID        uint64    `json:"id,omitempty"`
	HostID    uint64    `sql:"host_id" json:"host"`
	Name      string    `json:"name" binding:"required"`
	Data      JSON      `json:"data" binding:"required"`
	CreatedAt time.Time `sql:"default:now()" json:"createdAt,omitempty"`
}

// BeforeInsert hooks before add a new source
func (z *Source) BeforeInsert(db orm.DB) error {
	if z.CreatedAt.IsZero() {
		z.CreatedAt = time.Now()
	}
	return nil
}

// BeforeInsert hooks before add a new workspace
func (z *Workspace) BeforeInsert(db orm.DB) error {
	if z.CreatedAt.IsZero() {
		z.CreatedAt = time.Now()
	}
	return nil
}

// BeforeInsert hooks before add a new host
func (z *Host) BeforeInsert(db orm.DB) error {
	if z.CreatedAt.IsZero() {
		z.CreatedAt = time.Now()
	}
	return nil
}

// BeforeInsert hooks before add a new service
func (z *Service) BeforeInsert(db orm.DB) error {
	if z.CreatedAt.IsZero() {
		z.CreatedAt = time.Now()
	}
	return nil
}

// BeforeInsert hooks before add a new information
func (z *Info) BeforeInsert(db orm.DB) error {
	if z.CreatedAt.IsZero() {
		z.CreatedAt = time.Now()
	}
	return nil
}
