/* For license and copyright info please see LEGAL file in repository */

package error

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

// NewChain wrap an error and additional information usually use in logging to save more details about error.
func NewChain(err *Error, info string) (ce *ChainError) {
	if err == nil {
		return
	}
	return &ChainError{
		Err: err,
		info:  info,
	}
}

// ChainError is a extended implementation of Error to carry custom details along error.
type ChainError struct {
	*Err
	info string
}

func (ce *ChainError) PastChain() protocol.Error { return ce.Err }
func (ce *ChainError) ToString() string {
	return "\n" + ce.Err.ToString() + "\n	Chain Info: " + ce.info
}

// Go compatibility methods. Unwrap provides compatibility for Go 1.13 error chains.
func (ce *ChainError) Error() string { return ce.ToString() }
func (ce *ChainError) Cause() error  { return ce.Err }
func (ce *ChainError) Unwrap() error { return ce.Err }
// func (ce *ChainError) Is(error) bool
// func (ce *ChainError) As(any) bool 
