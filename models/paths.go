package models

import "strconv"

// HTTPRoute is a route used in HTTP routes
type HTTPRoute string

func (r HTTPRoute) String() string {
	return string(r)
}

// ID return the path with the "id" param.
func (r HTTPRoute) ID() string {
	return string(r) + "/:id"
}

// WithID returns the same path, but with an ID (used to GET, PUT or DELETE resources)
func (r HTTPRoute) WithID(id uint64) string {
	return string(r) + "/" + strconv.FormatUint(id, 10)
}

// WorkspacePath is the HTTP path to manage workspaces
const WorkspacePath = HTTPRoute("/workspaces")

// SourcePath is the HTTP path to manage sources
const SourcePath = HTTPRoute("/sources")

// HostPath is the HTTP path to manage hosts
const HostPath = HTTPRoute("/hosts")

// ServicePath is the HTTP path to manage services
const ServicePath = HTTPRoute("/services")

// NotePath is the HTTP path to manage notes
const NotePath = HTTPRoute("/notes")

// IssuePath is the HTTP path to manage issues
const IssuePath = HTTPRoute("/issues")

// ImportPath is the HTTP path to import data
const ImportPath = "/import"
