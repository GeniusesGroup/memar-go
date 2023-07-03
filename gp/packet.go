/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"libgo/protocol"
)

type Frame []byte

// CheckFrame will check frame for any bad situation.
// Always check frame before use any other Frame methods otherwise Go panic occur.
func (f Frame) CheckFrame() protocol.Error {
	if len(f) != FrameLen {
		return &ErrFrameLength
	}
	if f.FrameID() != protocol.Network_FrameID_GP {
		return &ErrBadFrameID
	}
	return nil
}

func (f Frame) FrameID() protocol.Network_FrameID { return protocol.Network_FrameID(f[0]) }
func (f Frame) DestinationAddr() (addr Addr)      { copy(addr[:], f[1:]); return }
func (f Frame) SourceAddr() (addr Addr)           { copy(addr[:], f[17:]); return }

func (f Frame) SetFrameID(fID protocol.Network_FrameID) { f[0] = byte(fID) }
func (f Frame) SetDestinationAddr(addr Addr)            { copy(f[1:], addr[:]) }
func (f Frame) SetSourceAddr(addr Addr)                 { copy(f[17:], addr[:]) }

//libgo:impl libgo/protocol.Network_Frame
func (f Frame) FrameLen() (frameLength int) { return FrameLen }
func (f Frame) NextFrame() []byte           { return f[FrameLen:] }
func (f Frame) Process(soc protocol.Socket) (err protocol.Error) {
	var gpAddr Addr = f.SourceAddr()
	// Find Connection from ConnectionPoolByPeerAdd by requester GP
	var conn *Connection
	conn, err = conns.GetConnectionByPeerAddr(gpAddr)
	// If it is first time that user want to connect or longer than server GC old unused connections.
	if conn == nil {
		// TODO:::
		// conn, err = MakeNewConnectionByPeerAdd(gpAddr, appMux.nl)
		if err != nil {
			// Send response or just ignore frame
			conn.FailedPacketsReceived()
			// TODO::: DDOS!!??
			return
		}
		// appMux.Connections.RegisterConnection(conn)
	}

	// Metrics data
	// conn.PacketReceived(uint64(len(frame)))
	// conn.PacketPayloadSize = GetPayloadLength(frame) // It's not working due to frame not encrypted yet.

	return
}
func (f Frame) Do(soc protocol.Socket) (err protocol.Error) {
	return
}
