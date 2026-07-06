/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	error_p "memar/process/error/protocol"
)

// Encoder store data to encode given data by each method!
type EncoderOptions struct {
	// quoted causes primitive fields to be encoded inside JSON strings.
	Quoted bool
	// escapeHTML causes '<', '>', and '&' to be escaped in JSON strings.
	EscapeHTML bool
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (eo *EncoderOptions) Init() (err error_p.Error)   { return }
func (eo *EncoderOptions) Reinit() (err error_p.Error) { return }
func (eo *EncoderOptions) Deinit() (err error_p.Error) { return }

func (eo *EncoderOptions) SetDefaults() {}
