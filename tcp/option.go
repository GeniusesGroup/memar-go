/* For license and copyright information please see LEGAL file in repository */

package tcp

type Option struct {
	OptionKind   OptionKind
	OptionLength uint8
	OptionData   []byte
}

// OptionKind represents a TCP option kind code.
type OptionKind uint8

const (
	OptionKind_EndList                         = 0
	OptionKind_Nop                             = 1
	OptionKind_MSS                             = 2  // len = 4
	OptionKind_WindowScale                     = 3  // len = 3
	OptionKind_SACKPermitted                   = 4  // len = 2
	OptionKind_SACK                            = 5  // len = n
	OptionKind_Echo                            = 6  // len = 6, obsolete
	OptionKind_EchoReply                       = 7  // len = 6, obsolete
	OptionKind_Timestamps                      = 8  // len = 10
	OptionKind_PartialOrderConnectionPermitted = 9  // len = 2, obsolete
	OptionKind_PartialOrderServiceProfile      = 10 // len = 3, obsolete
	OptionKind_CC                              = 11 // obsolete
	OptionKind_CCNew                           = 12 // obsolete
	OptionKind_CCEcho                          = 13 // obsolete
	OptionKind_AltChecksum                     = 14 // len = 3, obsolete
	OptionKind_AltChecksumData                 = 15 // len = n, obsolete
)
