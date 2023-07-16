/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

/*
type option struct {
	Kind    byte
	Payload []byte // Can be nil in some kinds
}
*/
type Options []byte

func (o Options) Kind() optionKind { return optionKind(o[0]) }
func (o Options) Payload() []byte  { return o[1:] }

// OptionKind represents a TCP option kind code.
type optionKind byte

// https://www.iana.org/assignments/tcp-parameters/tcp-parameters.xhtml
// https://datatracker.ietf.org/doc/html/rfc4413#section-4.3.1
const (
	OptionKind_EndList optionKind = iota
	OptionKind_Nop
	OptionKind_MSS                             // len = 4, Maximum Segment Size
	OptionKind_WindowScale                     // len = 3
	OptionKind_SACKPermitted                   // len = 2
	OptionKind_SACK                            // len = n
	OptionKind_Echo                            // len = 6, obsolete
	OptionKind_EchoReply                       // len = 6, obsolete
	OptionKind_Timestamps                      // len = 10
	OptionKind_PartialOrderConnectionPermitted // len = 2, obsolete
	OptionKind_PartialOrderServiceProfile      // len = 3, obsolete
	OptionKind_CC                              // obsolete
	OptionKind_CCNew                           // obsolete
	OptionKind_CCEcho                          // obsolete
	OptionKind_AltChecksum                     // len = 3, obsolete
	OptionKind_AltChecksumData                 // len = n, obsolete
)
