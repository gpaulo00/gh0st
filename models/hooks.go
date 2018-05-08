package models

import (
	"time"

	"github.com/go-pg/pg/orm"
)

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

// BeforeInsert hooks before add a new note
func (z *Note) BeforeInsert(db orm.DB) error {
	if z.CreatedAt.IsZero() {
		z.CreatedAt = time.Now()
	}
	return nil
}

// BeforeInsert hooks before add a new issue
func (z *Issue) BeforeInsert(db orm.DB) error {
	if z.CreatedAt.IsZero() {
		z.CreatedAt = time.Now()
	}
	return nil
}
