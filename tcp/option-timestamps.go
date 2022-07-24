/* For license and copyright information please see LEGAL file in repository */

package tcp

import "github.com/GeniusesGroup/libgo/binary"

type optionTimestamps []byte

func (o optionTimestamps) Length() byte       { return o[0] }
func (o optionTimestamps) Timestamps() uint16 { return binary.BigEndian.Uint16(o[1:]) }
func (o optionTimestamps) NextOption() []byte { return o[8:] }

func (o optionTimestamps) Process(s *Socket) error {
	return nil
}
