/* For license and copyright info please see LEGAL file in repository */

package error

import (
	"../protocol"
)

// NewChain wrap an error and additional information usually use in logging to save more details about error.
func NewChain(err protocol.Error, info string) (ce *ChainError) {
	if err == nil {
		return
	}
	return &ChainError{
		err:  err,
		info: info,
	}
}

// ChainError is a extended implementation of Error to carry custom details along error.
type ChainError struct {
	err  protocol.Error
	info string
}

func (ce *ChainError) PastChain() *ChainError        { return ce.err.(*ChainError) }
func (ce *ChainError) ID() uint64                    { return ce.err.ID() }
func (ce *ChainError) MediaType() protocol.MediaType { return ce.err.MediaType() }
func (ce *ChainError) ToString() string {
	return "\n	Chain Error - Cause by: " + ce.err.ToString() + " - Info: " + ce.info
}
func (ce *ChainError) Equal(err protocol.Error) bool { return ce.err.Equal(err) }

// Unwrap provides compatibility for Go 1.13 error chains.
func (ce *ChainError) Error() string {
	return "\n	Chain Error - Cause: " + ce.err.Error() + " - Info: " + ce.info
}
func (ce *ChainError) Cause() error  { return ce.err }
func (ce *ChainError) Unwrap() error { return ce.err }
