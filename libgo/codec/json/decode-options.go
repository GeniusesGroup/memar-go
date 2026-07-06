/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	error_p "memar/process/error/protocol"
)

// Decoder store data to decode data by each method.
type DecodeOptions struct {
	ErrorOnDuplicateKey bool
	ErrorOnNotFoundKey  bool
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (do *DecodeOptions) Init() (err error_p.Error)   { return }
func (do *DecodeOptions) Reinit() (err error_p.Error) { return }
func (do *DecodeOptions) Deinit() (err error_p.Error) { return }

func (do *DecodeOptions) SetDefaults() {}
