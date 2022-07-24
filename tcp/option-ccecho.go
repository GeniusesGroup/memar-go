/* For license and copyright information please see LEGAL file in repository */

package tcp

import "github.com/GeniusesGroup/libgo/binary"

type optionCCEcho []byte

func (o optionCCEcho) Length() byte       { return o[0] }
func (o optionCCEcho) CCEcho() uint16     { return binary.BigEndian.Uint16(o[1:]) }
func (o optionCCEcho) NextOption() []byte { return o[5:] }

func (o optionCCEcho) Process(s *Socket) error {
	return nil
}
