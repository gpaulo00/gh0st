package custom

import (
	"encoding/json"
	"errors"
	"io"
	"sync"

	"github.com/go-pg/pg"
	"github.com/gpaulo00/gh0st/models"
)

// Import inserts all the data into database
func (s *ImportSource) Import(tx *pg.Tx) error {
	hosts := make([]interface{}, len(s.Hosts))
	var (
		services []interface{}
		notes    []interface{}
		issues   []interface{}
	)

	// add source
	if err := tx.Insert(&s.Source); err != nil {
		return err
	}

	// without hosts?
	if len(s.Hosts) <= 0 {
		return errors.New("the source is not adding information")
	}

	// add hosts
	for i := range s.Hosts {
		host := &s.Hosts[i].Host
		host.SourceID = s.Source.ID
		hosts[i] = host
	}
	if _, err := tx.Model(hosts...).Insert(); err != nil {
		return err
	}

	// parse services, notes & issues
	for i := range s.Hosts {
		hostID := hosts[i].(*models.Host).ID
		var wg sync.WaitGroup

		// parse services
		go func(i []models.Service) {
			wg.Add(1)
			defer wg.Done()
			for _, srv := range i {
				srv.HostID = hostID
				services = append(services, &srv)
			}
		}(s.Hosts[i].Services)

		// parse notes
		go func(i []models.Note) {
			wg.Add(1)
			defer wg.Done()
			for _, note := range i {
				note.HostID = hostID
				notes = append(notes, &note)
			}
		}(s.Hosts[i].Notes)

		// parse issues
		go func(i []models.Issue) {
			wg.Add(1)
			defer wg.Done()
			for _, issue := range i {
				issue.HostID = hostID
				issues = append(issues, &issue)
			}
		}(s.Hosts[i].Issues)

		wg.Wait()
	}

	// insert all
	if err := tx.Insert(services...); err != nil {
		return err
	}
	if err := tx.Insert(notes...); err != nil {
		return err
	}
	return tx.Insert(issues...)
}

// Integration defines how to import a raw source into gh0st
type Integration struct {
	db *pg.Tx
}

// Parse makes a *Root element from an io.Reader
func (n *Integration) Parse(r io.Reader) error {
	res := new(ImportSource)
	if err := json.NewDecoder(r).Decode(res); err != nil {
		return err
	}

	return res.Import(n.db)
}

// New returns the Custom plugin.
func New(db *pg.Tx) *Integration {
	return &Integration{db}
}
