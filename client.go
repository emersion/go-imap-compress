package compress

import (
	"compress/flate"
	"errors"
	"net"

	imapclient "github.com/emersion/go-imap/client"
)

type Client struct {
	client *imapclient.Client
	isCompressed bool
}

// Instructs the server to use the named compression mechanism for all commands
// and/or responses.
func (c *Client) Compress(mech string) (err error) {
	if c.isCompressed {
		err = errors.New("COMPRESS is already enabled")
		return
	}

	if mech != Deflate {
		err = errors.New("Cannot start compression: mechanism " + mech + " not supported")
		return
	}

	cmd := &Command{Mechanism: mech}

	err = c.client.Upgrade(func (conn net.Conn) (net.Conn, error) {
		status, err := c.client.Execute(cmd, nil)
		if err != nil {
			return nil, err
		}
		if err = status.Err(); err != nil {
			return nil, err
		}

		return NewDeflateConn(conn, flate.DefaultCompression)
	})
	if err != nil {
		return
	}

	c.isCompressed = true
	return
}

// Check if this client's connection is compressed.
func (c *Client) IsCompressed() bool {
	return c.isCompressed
}

// Check if the server supports a compression mechanism.
func (c *Client) SupportsCompression(mech string) bool {
	return c.client.Caps[Capability + "=" + mech]
}

// Create a new client.
func NewClient(c *imapclient.Client) *Client {
	return &Client{client: c}
}
