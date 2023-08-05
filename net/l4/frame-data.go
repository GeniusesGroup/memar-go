/* For license and copyright information please see the LEGAL file in the code repository */

package l4

import (
	"memar/binary"
	"memar/protocol"
)

/*
	type DataFrame struct {
		Length   [2]byte // including the header fields
		StreamID [8]byte // uint64
		Offset   [8]byte // uint32
		Payload  []byte
	}
*/
type DataFrame []byte

func (f DataFrame) Length() uint16   { return binary.BigEndian(f[0:]).Uint16() }
func (f DataFrame) StreamID() uint64 { return binary.BigEndian(f[2:]).Uint64() }
func (f DataFrame) Offset() uint32   { return binary.BigEndian(f[10:]).Uint32() }
func (f DataFrame) Payload() []byte  { return f[18:f.Length()] }

//memar:impl memar/protocol.Network_Frame
func (f DataFrame) NextFrame() []byte { return f[f.Length():] }

func (f DataFrame) Process(sk protocol.Socket) (err protocol.Error) {
	// TODO::: check socket situation first

	// add payload to the requested offset of the stream payload
	var payloadCompleted = f.addDataToStream(sk)
	if payloadCompleted {
		sk.ScheduleProcessingSocket()
	}
	return
}

func (f DataFrame) Do(sk protocol.Socket) (err protocol.Error) {
	// TODO:::
	return
}

func (f DataFrame) addDataToStream(sk protocol.Socket) (payloadCompleted bool) {
	// var offset = f.Offset()
	// var payload = f.Payload()
	// TODO:::
	return
}

/*
func handleDataFrame(sk protocol.Socket, DataFrame []byte, packetID uint32) (err protocol.Error) {
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
