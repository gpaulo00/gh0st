package nmap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"sync"
	"text/template"
	"time"

	"github.com/go-pg/pg"
	"github.com/gobuffalo/packr"
	"github.com/gpaulo00/gh0st/models"
)

var ourTemplate *template.Template

func init() {
	var err error
	box := packr.NewBox("./templates")
	ourTemplate, err = template.New("nmap").Parse(box.String("issue.md"))
	if err != nil {
		panic(err)
	}
}

// Import inserts all the Nmap data into database
func (r *Root) Import(db *pg.DB, ws uint64) error {
	at := time.Time(r.Start)
	used := map[int]int{} // used hosts indexes
	var (
		hosts    []interface{} // *models.Host
		services []interface{} // *models.Service
		issues   []interface{} // *models.Issue
	)

	err := db.RunInTransaction(func(tx *pg.Tx) error {
		// generate source entry
		source := models.Source{
			WorkspaceID: ws,
			Generator:   fmt.Sprintf("%s %s", r.Scanner, r.Version),
			GeneratedAt: &at,
			SourceInfo: &models.JSON{
				"arguments": r.Args,
				"type":      r.ScanInfo.Type,
				"verbose":   r.Verbose.Level,
				"debug":     r.Debugging.Level,
			},
		}
		if err := tx.Insert(&source); err != nil {
			return err
		}

		// import hosts
		for i, host := range r.Hosts {
			if len(host.Addresses) < 1 {
				continue
			}

			// parse address
			var ip net.IP
			for _, addr := range host.Addresses {
				ip = net.ParseIP(addr.Addr)
				if ip != nil {
					break
				}
			}
			if ip == nil {
				continue
			}
			hosts = append(hosts, &models.Host{
				SourceID: source.ID,
				Address:  ip,
				State:    host.Status.State,
			})
			used[i] = len(hosts) - 1
		}

		// abort if no hosts
		if len(hosts) < 1 {
			return nil
		}
		if _, err := tx.Model(hosts...).Insert(); err != nil {
			return err
		}

		// parse services and issues
		for i := range r.Hosts {
			idx, ok := used[i]
			if !ok {
				// skip
				continue
			}
			host := hosts[idx].(*models.Host)
			var wg sync.WaitGroup
			wg.Add(2)

			// services
			go func() {
				defer wg.Done()

				for _, port := range r.Hosts[i].Ports {
					service := port.Service.String()
					services = append(services, &models.Service{
						HostID:   host.ID,
						Protocol: port.Protocol,
						Port:     port.PortID,
						State:    port.State.State,
						Service:  &service,
					})
				}
			}()

			// generate issue of the host
			go func(host Host, mhost *models.Host) {
				defer wg.Done()

				buf := new(bytes.Buffer)
				if err := ourTemplate.Execute(buf, host); err != nil {
					panic(err)
				}
				issues = append(issues, &models.Issue{
					HostID:  mhost.ID,
					Level:   models.InfoIssue,
					Title:   "Nmap Scan",
					Content: buf.String(),
				})
			}(r.Hosts[i], host)
			wg.Wait()
		}

		// insert all
		if len(services) > 1 {
			if err := tx.Insert(services...); err != nil {
				return err
			}
		}
		return tx.Insert(issues...)
	})

	return err
}

// Integration defines how to import a Nmap scan into gh0st
type Integration struct {
	db *pg.DB
	ws uint64
}

// Parse makes a *Root element from an io.Reader
func (n *Integration) Parse(r io.Reader) error {
	res := new(Root)
	if err := xml.NewDecoder(r).Decode(res); err != nil {
		return err
	}

	return res.Import(n.db, n.ws)
}

// New returns the Nmap plugin.
func New(db *pg.DB, ws uint64) *Integration {
	return &Integration{db, ws}
}
