package plugins

import "io"

// Plugin is a module that parses data and imports it
// into the database
type Plugin interface {
	Parse(io.Reader) error
}
