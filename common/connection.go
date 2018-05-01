package common

import "net"

// IConnection describes a connection via gh0st
type IConnection interface {
	Connection() net.Conn
	Run() error
	Deattach() error
	Attach() error
	Close() error
}
