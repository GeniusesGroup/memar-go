/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import "github.com/GeniusesGroup/libgo/binary"

type optionWindowScale []byte

func (o optionWindowScale) Length() byte        { return o[0] }
func (o optionWindowScale) WindowScale() uint16 { return binary.BigEndian.Uint16(o[1:]) }
func (o optionWindowScale) NextOption() []byte  { return o[2:] }

// handler options -> socket
func (o optionWindowScale) Process(s *Socket) error {
	return nil
}
