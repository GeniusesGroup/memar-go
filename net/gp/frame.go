/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	container_p "memar/computer/adt/container/protocol"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
)

type Frame []byte

// CheckFrame will check frame for any bad situation.
// Always check frame before use any other Frame methods otherwise Go panic occur.
func (f Frame) CheckFrame() error_p.Error {
	if len(f) != FrameLen {
		return &ErrFrameLength
	}
	if f.FrameType() != net_p.FrameType_GP {
		return &ErrBadFrameType
	}
	return nil
}

//memar:impl memar/net/protocol.Field_FrameType
func (f Frame) FrameType() net_p.FrameType { return net_p.FrameType(f[0]) }

//memar:impl memar/net/protocol.Field_NetworkAddresses
func (f Frame) DestinationAddr() (addr Addr) { copy(addr[:], f[1:]); return }
func (f Frame) SourceAddr() (addr Addr)      { copy(addr[:], f[17:]); return }

func (f Frame) setFrameType(fID net_p.FrameType) { f[0] = byte(fID) }
func (f Frame) SetDestinationAddr(addr Addr)     { copy(f[1:], addr[:]) }
func (f Frame) SetSourceAddr(addr Addr)          { copy(f[17:], addr[:]) }

//memar:impl memar/protocol.Network_Frame
func (f Frame) FrameLength() (frameLength container_p.NumberOfElement) { return FrameLen }
func (f Frame) NextFrame() []byte                                      { return f[FrameLen:] }
func (f Frame) Process(soc net_p.Socket) (err error_p.Error) {
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
			// conn.FailedPacketsReceived()
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
func (f Frame) Do(soc net_p.Socket) (err error_p.Error) {
	return
}
