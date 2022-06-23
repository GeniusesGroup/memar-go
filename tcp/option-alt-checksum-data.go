/* For license and copyright information please see LEGAL file in repository */

package tcp

import "../binary"

type optionAltChecksumData []byte

func (o optionAltChecksumData) Length() byte            { return o[0] }
func (o optionAltChecksumData) AltChecksumData() uint16 { return binary.BigEndian.Uint16(o[1:]) } // unrecognised
func (o optionAltChecksumData) NextOption() []byte      { return o[3:] }

func (o optionAltChecksumData) Process(s *Socket) error {
	return nil
}
