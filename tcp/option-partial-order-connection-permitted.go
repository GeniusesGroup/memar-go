/* For license and copyright information please see LEGAL file in repository */

package tcp

type optionPartialOrderConnectionPermitted []byte

func (o optionPartialOrderConnectionPermitted) Length() byte { return o[0] }

// func (o optionPartialOrderConnectionPermitted) PartialOrderConnectionPermitted() uint16 {
// 	return binary.BigEndian.Uint16(o[1:])
// }
func (o optionPartialOrderConnectionPermitted) NextOption() []byte { return o[3:] }

func (o optionPartialOrderConnectionPermitted) Process(s *Socket) error {
	return nil
}
