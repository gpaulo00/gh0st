package ncat

import (
	"context"
	"errors"
	"io"
	"net"
	"os"

	"bitbucket.org/gpaulo00/gh0st/common"
)

// Connection describes a Ncat connection, and implements IConnection.
type Connection struct {
	Addr   *net.TCPAddr
	Conn   net.Conn
	cancel context.CancelFunc
}

// Run executes the connection to a Ncat server
func (conn *Connection) Run() error {
	c, err := net.DialTCP("tcp", nil, conn.Addr)
	if err != nil {
		return err
	}

	// all right
	conn.Conn = c
	return nil
}

// Connection returns the wrapped net.Conn
func (conn *Connection) Connection() net.Conn {
	return conn.Conn
}

// Close closes the wrapped net.Conn and deattaches the I/O
func (conn *Connection) Close() error {
	// deattach
	conn.Deattach()

	// close connection
	if conn.Conn != nil {
		return conn.Conn.Close()
	}
	return nil
}

// Deattach the connection I/O.
func (conn *Connection) Deattach() error {
	if conn.cancel == nil {
		return errors.New("Connection is not attached")
	}

	conn.cancel()
	conn.cancel = nil
	return nil
}

// Attach binds the connection I/O to standard I/O.
func (conn *Connection) Attach() error {
	if conn.cancel != nil {
		return errors.New("Connection is already attached")
	}
	ctx, cancel := context.WithCancel(context.Background())
	go io.Copy(os.Stdout, common.NewReader(ctx, conn.Conn))
	go io.Copy(common.NewWriter(ctx, conn.Conn), os.Stdin)
	conn.cancel = cancel
	return nil
}

// NewConnection creates a new Connection value
func NewConnection(addr *net.TCPAddr) *Connection {
	return &Connection{Addr: addr}
}
