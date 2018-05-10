package nmap

// Based on https://github.com/lair-framework/go-nmap/blob/master/nmap.go (MIT License)
// With some modifications

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
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

	// extra
	ProfileName string `xml:"profile_name,attr"`

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

	// extra
	TCPSequence  TCPSequence  `xml:"tcpsequence"`
	IPIDSequence IPIDSequence `xml:"ipidsequence"`
	HostScripts  []Script     `xml:"hostscript>script"`
}

// TCPSequence contains information regarding the detected tcp sequence.
type TCPSequence struct {
	Index      int    `xml:"index,attr"`
	Difficulty string `xml:"difficulty,attr"`
	Values     string `xml:"values,attr"`
}

// Sequence contains information regarding the detected X sequence.
type Sequence struct {
	Class  string `xml:"class,attr"`
	Values string `xml:"values,attr"`
}

// IPIDSequence contains info about the ipid sequence
type IPIDSequence Sequence

// Script contains information from Nmap Scripting Engine.
type Script struct {
	ID     string  `xml:"id,attr"`
	Output string  `xml:"output,attr"`
	Tables []Table `xml:"table"`
}

// Table contains the output of the script in a more parse-able form.
// ToDo: This should be a map[string][]string
type Table struct {
	Key      string   `xml:"key,attr"`
	Elements []string `xml:"elem"`
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

func (s Service) String() string {
	if s.Product == "" && s.Version == "" {
		return "unknown"
	}
	return fmt.Sprintf("%s %s", s.Product, s.Version)
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
	Name     string `xml:"name,attr"`
	Accuracy string `xml:"accuracy,attr"`
	Line     string `xml:"line,attr" `
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
	Proto string `xml:"proto,attr"`
	Port  int    `xml:"port,attr"`
	Hops  []Hop  `xml:"hop"`
}

// Hop is a ip hop to a Host.
type Hop struct {
	TTL    float32 `xml:"ttl,attr"`
	RTT    float32 `xml:"rtt,attr"`
	IPAddr string  `xml:"ipaddr,attr"`
	Host   string  `xml:"host,attr"`
}
