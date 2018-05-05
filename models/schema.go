package models

import (
	"net"
	"time"

	"github.com/go-pg/pg/orm"
)

// Workspace organizes all the data
type Workspace struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `sql:"default:now()" json:"createdAt"`
}

// Source is the source of data
type Source struct {
	ID          uint64    `json:"id"`
	WorkspaceID uint64    `sql:"workspace_id" json:"workspace" binding:"required"`
	Generator   string    `json:"generator" binding:"required"`
	CreatedAt   time.Time `sql:"default:now()" json:"createdAt"`
}

// Host is a target of data
type Host struct {
	ID        uint64    `json:"id"`
	SourceID  uint64    `sql:"source_id" json:"source"`
	Address   net.IP    `json:"address" binding:"required"`
	CreatedAt time.Time `sql:"default:now()" json:"createdAt"`
}

// Service is running service of a host
type Service struct {
	ID        uint64    `json:"id"`
	HostID    uint64    `sql:"host_id" json:"host"`
	Port      uint16    `json:"port" binding:"required"`
	Service   string    `json:"service" binding:"required"`
	CreatedAt time.Time `sql:"default:now()" json:"createdAt"`
}

// Info is an extra information of a host
type Info struct {
	ID        uint64                 `json:"id"`
	HostID    uint64                 `sql:"host_id" json:"host"`
	Name      string                 `json:"name" binding:"required"`
	Data      map[string]interface{} `json:"data" binding:"required"`
	CreatedAt time.Time              `sql:"default:now()" json:"createdAt"`
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
