# go-imap-compress

[![GoDoc](https://godoc.org/github.com/emersion/go-imap-compress?status.svg)](https://godoc.org/github.com/emersion/go-imap-compress)

The [IMAP COMPRESS Extension](https://tools.ietf.org/html/rfc4978) for [go-imap](https://github.com/emersion/go-imap)

## Usage

```go
package main

import (
		"log"

		"github.com/emersion/go-imap/client"
		"github.com/emersion/go-imap-compress"
)

func main() {
	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS("mail.example.org:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login("username", "password"); err != nil {
		log.Fatal(err)
	}

	// Get capabilities if needed
	if c.Caps == nil {
		if _, err := c.Capability(); err != nil {
			log.Fatal(err)
		}
	}

	// Enable compression if possible
	comp := compress.NewClient(c)
	if comp.SupportsCompression(compress.Deflate) {
		if err := comp.Compress(compress.Deflate); err != nil {
			log.Fatal(err)
		}
	}

	// Continue using c with compression enabled
}
```

## License

MIT
