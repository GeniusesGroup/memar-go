/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"../protocol"
	"../syllab"
)

/*
type dataFrame struct {
	Length   [2]byte // including the header fields
	StreamID [4]byte // uint32
	Offset   [4]byte // uint32 also can act as offset
	Payload  []byte
}
*/
type dataFrame []byte

func (f dataFrame) Length() uint16    { return syllab.GetUInt16(f, 0) }
func (f dataFrame) StreamID() uint32  { return syllab.GetUInt32(f, 2) }
func (f dataFrame) Offset() uint32    { return syllab.GetUInt32(f, 6) }
func (f dataFrame) Payload() []byte   { return f[10:f.Length()] }
func (f dataFrame) NextFrame() []byte { return f[f.Length():] }

// appendData add data to the requested offset of the stream
func appendData(conn protocol.Connection, frame dataFrame) (err protocol.Error) {
	var streamID uint32 = frame.StreamID()
	var stream protocol.Stream
	stream, err = conn.Stream(streamID)
	if err != nil {
		conn.StreamFailed()
		// Send response or just ignore stream
		// TODO::: DDOS!!??
		return
	}

	// TODO:::
	// add payload to Stream payload
	// var offset uint32 = frame.Offset()
	// err = addNewGPPacket(stream, GetPayload(packet), packetID)

	if stream.Status() == protocol.ConnectionStateReady {
		// decide by stream odd or even
		// TODO::: check better performance as streamID&1 to check odd id
		if streamID%2 == 0 {
			err = stream.Protocol().HandleIncomeRequest(stream)
			if err == nil {
				conn.StreamSucceed()
			} else {
				conn.StreamFailed()
			}
		} else {
			// income response
			stream.SetState(protocol.ConnectionStateReady)
		}
	}
	return
}
