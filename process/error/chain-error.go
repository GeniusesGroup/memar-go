/* For license and copyright info please see LEGAL file in repository */

package error

import (
	error_p "memar/error/protocol"
)

// PastChain unwrap an error and return last chain of error.
func PastChain(err error_p.Error) (last error_p.Error) {
	var chainedError, ok = err.(error_p.Chain)
	if ok {
		return chainedError.PastChain()
	}
	return
}

// NewChain wrap an error and additional information usually use in logging to save more details about error.
func NewChain(err error_p.Error, chain error_p.Error) (ce *ChainError) {
	if err == nil || chain == nil {
		return
	}
	return &ChainError{
		Error: err,
		chain: chain,
	}
}

// ChainError is a extended implementation of Error to carry custom details along error.
type ChainError struct {
	error_p.Error
	chain error_p.Error
}

func (ce *ChainError) PastChain() error_p.Error { return ce.chain }
