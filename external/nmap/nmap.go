package nmap

// Based on https://github.com/lair-framework/go-nmap/blob/master/nmap.go (MIT License)
// With some modifications

import (
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"

	"github.com/gpaulo00/gh0st/models"
)

// Timestamp is a parsed time.Time from xml
type Timestamp time.Time

// str2time converts a string containing a UNIX timestamp to to a time.Time.
func (t *Timestamp) str2time(s string) (err error) {
	ts, err := strconv.Atoi(s)
	if err != nil {
		return
	}
	*t = Timestamp(time.Unix(int64(ts), 0))
	return
}

// time2str formats the time.Time value as a UNIX timestamp string.
// XXX these might also need to be changed to pointers. See str2time and UnmarshalXMLAttr.
func (t Timestamp) time2str() string {
	return fmt.Sprint(time.Time(t))
}

// MarshalXMLAttr marshals the Timestamp as a xml attribute
func (t *Timestamp) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: t.time2str()}, nil
}

// UnmarshalXMLAttr unmarshals the Timestamp from a xml attribute
func (t *Timestamp) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	return t.str2time(attr.Value)
}

// Root is contains all the data for a single nmap scan.
type Root struct {
	XMLName xml.Name `xml:"nmaprun"`
	// Generator
	Scanner string `xml:"scanner,attr"`
	Version string `xml:"version,attr"`

	// Generated At
	Start Timestamp `xml:"start,attr"`

	// source info
	Args      string    `xml:"args,attr"`
	ScanInfo  ScanInfo  `xml:"scaninfo"`
	Verbose   Verbose   `xml:"verbose"`
	Debugging Debugging `xml:"debugging"`

	// hosts
	Hosts []Host `xml:"host"`
}

// ScanInfo contains informational regarding how the scan
// was run.
type ScanInfo struct {
	Type        string `xml:"type,attr"`
	Protocol    string `xml:"protocol,attr"`
	NumServices int    `xml:"numservices,attr"`
	Services    string `xml:"services,attr"`
	ScanFlags   string `xml:"scanflags,attr"`
}

// Verbose contains the verbosity level for the Nmap scan.
type Verbose struct {
	Level int `xml:"level,attr"`
}

// Debugging contains the debugging level for the Nmap scan.
type Debugging struct {
	Level int `xml:"level,attr"`
}

// Host contains all information about a single host.
type Host struct {
	// host
	Addresses []Address `xml:"address"`

	// services
	Ports      []Port       `xml:"ports>port"`
	ExtraPorts []ExtraPorts `xml:"ports>extraports"`

	// info
	Comment   string     `xml:"comment,attr"`
	Status    Status     `xml:"status"`
	Hostnames []Hostname `xml:"hostnames>hostname"`
	Os        Os         `xml:"os"`
	Distance  Distance   `xml:"distance"`
	Uptime    Uptime     `xml:"uptime"`
	Trace     Trace      `xml:"trace"`
}

// Status is the host's status. Up, down, etc.
type Status struct {
	State     string  `xml:"state,attr"`
	Reason    string  `xml:"reason,attr"`
	ReasonTTL float32 `xml:"reason_ttl,attr"`
}

// Address contains a IPv4 or IPv6 address for a Host.
type Address struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
	Vendor   string `xml:"vendor,attr"`
}

// Hostname is a single name for a Host.
type Hostname struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

// ExtraPorts contains the information about the closed|filtered ports.
type ExtraPorts struct {
	State string `xml:"state,attr"`
	Count int    `xml:"count,attr"`
}

// Port contains all the information about a scanned port.
type Port struct {
	Protocol string  `xml:"protocol,attr"`
	PortID   uint16  `xml:"portid,attr"`
	State    State   `xml:"state"`
	Owner    Owner   `xml:"owner"`
	Service  Service `xml:"service"`
}

// State contains information about a given ports
// status. State will be open, closed, etc.
type State struct {
	State     string  `xml:"state,attr"`
	Reason    string  `xml:"reason,attr"`
	ReasonTTL float32 `xml:"reason_ttl,attr"`
	ReasonIP  string  `xml:"reason_ip,attr"`
}

// Owner contains the name of Port.Owner.
type Owner struct {
	Name string `xml:"name,attr"`
}

// Service contains detailed information about a Port's
// service details.
type Service struct {
	Name       string `xml:"name,attr"`
	Conf       int    `xml:"conf,attr"`
	Method     string `xml:"method,attr"`
	Version    string `xml:"version,attr"`
	Product    string `xml:"product,attr"`
	ExtraInfo  string `xml:"extrainfo,attr"`
	Tunnel     string `xml:"tunnel,attr"`
	Proto      string `xml:"proto,attr"`
	Rpcnum     string `xml:"rpcnum,attr"`
	Lowver     string `xml:"lowver,attr"`
	Highver    string `xml:"hiver,attr"`
	Hostname   string `xml:"hostname,attr"`
	OsType     string `xml:"ostype,attr"`
	DeviceType string `xml:"devicetype,attr"`
	ServiceFp  string `xml:"servicefp,attr"`
	CPEs       []CPE  `xml:"cpe"`
}

// CPE (Common Platform Enumeration) is a standardized way to name software
// applications, operating systems, and hardware platforms.
type CPE string

// Os contains the fingerprinted operating system for a Host.
type Os struct {
	PortsUsed      []PortUsed      `xml:"portused"`
	OsMatches      []OsMatch       `xml:"osmatch"`
	OsFingerprints []OsFingerprint `xml:"osfingerprint"`
}

// PortUsed is the port used to fingerprint a Os.
type PortUsed struct {
	State  string `xml:"state,attr"`
	Proto  string `xml:"proto,attr"`
	PortID int    `xml:"portid,attr"`
}

// OsMatch contains detailed information regarding a Os fingerprint.
type OsMatch struct {
	Name     string `xml:"name,attr" json:"name"`
	Accuracy string `xml:"accuracy,attr" json:"accuracy"`
	Line     string `xml:"line,attr" json:"-"`
}

// OsFingerprint is the actual fingerprint string.
type OsFingerprint struct {
	Fingerprint string `xml:"fingerprint,attr"`
}

// Distance is the amount of hops to a particular host.
type Distance struct {
	Value int `xml:"value,attr"`
}

// Uptime is the amount of time the host has been up.
type Uptime struct {
	Seconds  int    `xml:"seconds,attr"`
	Lastboot string `xml:"lastboot,attr"`
}

// Trace contains the hops to a Host.
type Trace struct {
	Proto string `xml:"proto,attr" json:"proto"`
	Port  int    `xml:"port,attr" json:"port"`
	Hops  []Hop  `xml:"hop" json:"hops"`
}

// Hop is a ip hop to a Host.
type Hop struct {
	TTL    float32 `xml:"ttl,attr" json:"ttl"`
	RTT    float32 `xml:"rtt,attr" json:"rtt"`
	IPAddr string  `xml:"ipaddr,attr" json:"address"`
	Host   string  `xml:"host,attr" json:"host"`
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
			srv := models.Service{
				Protocol: port.Protocol,
				Port:     port.PortID,
				State:    port.State.State,
				Service:  &port.Service.Name,
			}
			services = append(services, srv)
		}

		// infos
		info := models.Info{
			Name: "nmap data",
			Data: models.JSON{
				"comment":  host.Comment,
				"uptime":   host.Uptime.Seconds,
				"os":       host.Os.OsMatches,
				"distance": host.Distance.Value,
				"trace":    host.Trace,
			},
		}

		// append
		hosts = append(hosts, models.ImportHost{
			Host:     h,
			Services: services,
			Infos:    []models.Info{info},
		})
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
