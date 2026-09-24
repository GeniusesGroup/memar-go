/* For license and copyright information please see the LEGAL file in the code repository */

package l4

import (
	"memar/codec/binary"
	error_p "memar/process/error/protocol"
	net_p "memar/net/protocol"
)

/*
	type DataFrame struct {
		FrameLength [2]byte // including the header fields
		StreamID    [8]byte // uint64
		Offset      [8]byte // uint32
		Payload     []byte
	}
*/
type DataFrame []byte

func (f DataFrame) FrameLength() uint16 { return binary.BigEndian(f[0:]).Uint16() }
func (f DataFrame) StreamID() uint64    { return binary.BigEndian(f[2:]).Uint64() }
func (f DataFrame) Offset() uint32      { return binary.BigEndian(f[10:]).Uint32() }
func (f DataFrame) Payload() []byte     { return f[18:f.FrameLength()] }

//memar:impl memar/protocol.Network_Frame
func (f DataFrame) NextFrame() []byte { return f[f.FrameLength():] }

func (f DataFrame) Process(sk net_p.Socket) (err error_p.Error) {
	// TODO::: check socket situation first

	// add payload to the requested offset of the stream payload
	var payloadCompleted = f.addDataToStream(sk)
	if payloadCompleted {
		sk.OSI_ApplicationLayer().ScheduleProcessingSocket()
	}
	return
}

func (f DataFrame) Do(sk net_p.Socket) (err error_p.Error) {
	// TODO:::
	return
}

func (f DataFrame) addDataToStream(sk net_p.Socket) (payloadCompleted bool) {
	// var offset = f.Offset()
	// var payload = f.Payload()
	// TODO:::
	return
}

/*
func handleDataFrame(sk net_p.Socket, DataFrame []byte, packetID uint32) (err error_p.Error) {
	// Handle packet received not by order
	if packetID < st.LastPacketID {
		st.State = StateBrokenPacket
		err = &ErrPacketArrivedPosterior
	} else if packetID > st.LastPacketID+1 {
		st.State = StateBrokenPacket
		err = &ErrPacketArrivedAnterior
		// TODO::: send request to sender about not received packets!!
	} else if packetID+1 == st.LastPacketID {
		st.LastPacketID = packetID
	}
	// TODO::: non of above cover for packet 0||1 drop situation!

	// Use PacketID 0||1 for request||response to set stream settings!
	if packetID < 2 {
		// as.SetStreamSettings(st, p)
	} else {
		// TODO::: can't easily copy this way!!
		// copy(st.IncomePayload, p)
	}

	return
}
*/
