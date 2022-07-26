/* For license and copyright info please see LEGAL file in repository */

package error

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

// NewChain wrap an error and additional information usually use in logging to save more details about error.
func NewChain(err protocol.Error, info string) (ce *ChainError) {
	if err == nil {
		return
	}
	return &ChainError{
		Error: err,
		info:  info,
	}
}

// ChainError is a extended implementation of Error to carry custom details along error.
type ChainError struct {
	protocol.Error
	info string
}

func (ce *ChainError) PastChain() protocol.Error { return ce.Error }
func (ce *ChainError) ToString() string {
	return "\n" + ce.Error.ToString() + "\n	Chain Info: " + ce.info
}

// Go compatibility methods. Unwrap provides compatibility for Go 1.13 error chains.
// func (ce *ChainError) Error() string { return e.ToString() }
// func (ce *ChainError) Cause() error  { return ce }
// func (ce *ChainError) Unwrap() error { return ce }
