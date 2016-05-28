package compress

import (
	"errors"

	imap "github.com/emersion/go-imap/common"
)

// A COMPRESS command.
type Command struct {
	// Name of the compression mechanism.
	Mechanism string
}

func (cmd *Command) Command() *imap.Command {
	return &imap.Command{
		Name: CommandName,
		Arguments: []interface{}{cmd.Mechanism},
	}
}

func (cmd *Command) Parse(fields []interface{}) (err error) {
	if len(fields) < 1 {
		return errors.New("No enough arguments")
	}

	var ok bool
	if cmd.Mechanism, ok = fields[0].(string); !ok {
		return errors.New("Compression mechanism name must be a string")
	}

	return nil
}
