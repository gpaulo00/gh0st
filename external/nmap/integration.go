package nmap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"text/template"
	"time"

	"github.com/gpaulo00/gh0st/models"
)

const nmapIssue = `
# Nmap Issue
Issue generated by nmap module.

## Runnig Services
{{- if .Ports }}
| Port | State | Service |
| ---- | ----- | ------- |
{{- range $port := .Ports }}
| {{ $port.Protocol }}/{{ $port.PortID }} | {{ $port.State.State }} | {{ $port.Service }} |
{{- end }}
{{- else }}
Nmap could not find any open port in this host.
{{- end}}
`

const nmapNote = `
# Nmap Note
Note generated by nmap module.

## Operating System (OS)
{{- if .Os.OsMatches }}
| Name  | Accuracy |
| ----- | -------- |
{{- range $match := .Os.OsMatches }}
| {{ $match.Name }} | {{ $match.Accuracy }} |
{{- end }}
{{- else }}
Nmap could not guess the OS of this host.
{{- end }}

## Trace
| Address | Hostname | TTL | RTT |
| ------- | -------- | --- | --- |
{{- range $hop := .Trace.Hops }}
{{- $host := or $hop.Host "unknown" }}
| {{ $hop.IPAddr }} | {{ $host }} | {{ $hop.TTL }} | {{ $hop.RTT }}
{{- end }}

{{- if .Comment }}
## Comment
{{ .Comment }}
{{- end }}

{{- if ne .Uptime.Seconds 0 }}
## Uptime
{{ .Uptime.Lastboot }} ({{ .Uptime.Seconds }} sec.)
{{- end }}
`

var noteTemplate, issueTemplate *template.Template

func init() {
	var err error
	noteTemplate, err = template.New("nmap_note").Parse(nmapNote)
	if err != nil {
		panic(err)
	}

	issueTemplate, err = template.New("name_issue").Parse(nmapIssue)
	if err != nil {
		panic(err)
	}
}

// Import returns an gh0st-importable struct
func (r *Root) Import(ws uint64) *models.ImportForm {
	at := time.Time(r.Start)
	hosts := []models.ImportHost{}
	for _, host := range r.Hosts {
		if len(host.Addresses) < 1 {
			continue
		}

		// parse address
		ip := net.ParseIP(host.Addresses[0].Addr)
		if ip == nil {
			continue
		}
		h := models.Host{Address: ip, State: host.Status.State}

		// services
		services := []models.Service{}
		for _, port := range host.Ports {
			service := port.Service.String()
			srv := models.Service{
				Protocol: port.Protocol,
				Port:     port.PortID,
				State:    port.State.State,
				Service:  &service,
			}
			services = append(services, srv)
		}

		// generate note of the host
		buf := new(bytes.Buffer)
		if err := noteTemplate.Execute(buf, host); err != nil {
			panic(err)
		}
		note := models.Note{
			Title:   fmt.Sprintf("Nmap scan of %s", h.Address),
			Content: buf.String(),
		}

		// generate issue of the host
		single := models.ImportHost{
			Host:     h,
			Services: services,
			Notes:    []models.Note{note},
		}
		if len(host.Ports) > 0 {
			buf.Reset()
			if err := issueTemplate.Execute(buf, host); err != nil {
				panic(err)
			}
			single.Issues = []models.Issue{models.Issue{
				Title:   fmt.Sprintf("Nmap scan of %s", h.Address),
				Level:   string(models.LowIssue),
				Content: buf.String(),
			}}
		}

		// append
		hosts = append(hosts, single)
	}

	result := &models.ImportForm{
		Source: models.Source{
			WorkspaceID: ws,
			Generator:   fmt.Sprintf("%s %s", r.Scanner, r.Version),
			GeneratedAt: &at,
			SourceInfo: &models.JSON{
				"arguments": r.Args,
				"type":      r.ScanInfo.Type,
				"verbose":   r.Verbose.Level,
				"debug":     r.Debugging.Level,
			},
		},
		Hosts: hosts,
	}
	return result
}

// Integration defines how to import a Nmap scan into gh0st
type Integration struct{}

// Parse makes a *Root element from an io.Reader
func (n *Integration) Parse(r io.Reader) (*Root, error) {
	res := new(Root)
	if err := xml.NewDecoder(r).Decode(res); err != nil {
		return nil, err
	}
	return res, nil
}

// New returns a nmap integration api.
func New() *Integration {
	return &Integration{}
}