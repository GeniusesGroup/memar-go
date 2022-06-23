/* For license and copyright information please see LEGAL file in repository */

package tcp

type optionSACKPermitted []byte

func (o optionSACKPermitted) Length() byte { return o[0] }

// func (o optionSACKPermitted) SACKPermitted() uint16 { return binary.BigEndian.Uint16(o[1:]) }
func (o optionSACKPermitted) NextOption() []byte { return o[1:] }

func (o optionSACKPermitted) Process(s *Socket) error {
	return nil
}
