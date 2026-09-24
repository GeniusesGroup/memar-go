/* For license and copyright information please see the LEGAL file in the code repository */

package quic

import (
	"libgo/protocol"
)

// Indicate standard listen and send port number register for QUIC protocol.
const (
	ProtocolPortQUICReceive uint16 = 120
	ProtocolPortQUICSend    uint16 = 121
)

// QUICIncomeRequestHandler handle incoming QUIC request streams!
// Protocol Standard : https://datatracker.ietf.org/doc/html/rfc9000
// https://www.rfc-editor.org/rfc/rfc9000.html
func QUICIncomeRequestHandler(sk protocol.Socket) {

}

// https://github.com/quic-go/quic-go/
// https://github.com/goburrow/quic/blob/main/transport/frame.go
// https://github.com/lucas-clemente/quic-go/blob/master/framer.go
// https://github.com/microsoft/msquic
// https://github.com/cloudflare/quiche
