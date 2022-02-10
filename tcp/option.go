/* For license and copyright information please see LEGAL file in repository */

package tcp

type Options []byte

func (o Options) HasNext() bool {
	return len(o) > 0
}

func (o Options) Next() (opt Option, remain Options) {
	opt.Kind = OptionKind(o[0])
	switch opt.Kind {
	case OptionKind_EndList, OptionKind_Nop:
		opt.Length = 1
	default:
		opt.Length = o[1]
		opt.Data = o[2:opt.Length]
	}
	remain = o[opt.Length:]
	return
}

// https://datatracker.ietf.org/doc/html/rfc4413#section-4.3.1
type Option struct {
	Kind   OptionKind
	Length uint8 // including the header fields
	Data   []byte
}

// OptionKind represents a TCP option kind code.
type OptionKind uint8

// https://www.iana.org/assignments/tcp-parameters/tcp-parameters.xhtml
const (
	OptionKind_EndList OptionKind = iota
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
