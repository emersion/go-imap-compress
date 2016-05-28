package compress

import (
	"compress/flate"
	"io"
	"net"
)

type conn struct {
	net.Conn

	r io.ReadCloser
	w *flate.Writer
}

func (c *conn) Read(b []byte) (int, error) {
	return c.r.Read(b)
}

func (c *conn) Write(b []byte) (int, error) {
	return c.w.Write(b)
}

func (c *conn) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}

	if err := c.w.Close(); err != nil {
		return err
	}

	return c.Conn.Close()
}

func NewDeflateConn(c net.Conn, level int) (net.Conn, error) {
	r := flate.NewReader(c)
	w, err := flate.NewWriter(c, level)
	if err != nil {
		return nil, err
	}

	return &conn{
		Conn: c,
		r: r,
		w: w,
	}, nil
}