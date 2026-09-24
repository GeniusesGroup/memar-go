/* For license and copyright information please see the LEGAL file in the code repository */

package error

import (
	error_p "memar/process/error/protocol"
)

// PastChain unwrap an error and return last chain of error.
func PastChain(err error_p.Error) (last error_p.Error) {
	var chainedError, ok = err.(error_p.Chain)
	if ok {
		return chainedError.PastErrorChain()
	}
	return
}

// NewChain wrap an error and additional information usually use in logging to save more details about error.
func NewChain(err error_p.Error, chain error_p.Error) (self *ChainError) {
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

func (self *ChainError) PastChain() error_p.Error { return self.chain }
