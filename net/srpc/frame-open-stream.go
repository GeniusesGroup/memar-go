/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	"memar/codec/binary"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
	operation_p "memar/process/operation/protocol"
	"memar/process/services"
)

/*
	type OpenStreamFrame struct {
		HandlerID   uint64 // protocol ID usage is like TCP||UDP ports that indicate payload protocol.
		ServiceID   uint64
		CompressID  uint64
		DataLength  uint64
		Weight      operation_p.Weight
	}

TotalPacket uint32 // Expected packets count that send over this stream.
*/
type OpenStreamFrame []byte

func (self OpenStreamFrame) HandlerID() uint64          { return binary.LittleEndian(self[0:]).Uint64() }
func (self OpenStreamFrame) ServiceID() uint64          { return binary.LittleEndian(self[2:]).Uint64() }
func (self OpenStreamFrame) CompressID() uint64         { return binary.LittleEndian(self[10:]).Uint64() }
func (self OpenStreamFrame) DataLength() uint64         { return binary.LittleEndian(self[18:]).Uint64() }
func (self OpenStreamFrame) Weight() operation_p.Weight { return operation_p.Weight(self[24]) }

//memar:impl memar/protocol.Network_Frame
func (self OpenStreamFrame) NextFrame() []byte { return self[25:] }

func (self OpenStreamFrame) Do(sk net_p.Socket) (err error_p.Error) {
	// TODO::: allow multiple settings set??

	// Check server supported requested protocol
	var serviceID = operation_p.ID(self.ServiceID())
	_, err = services.GetByID(serviceID)
	if err != nil {
		// Send response or just ignore packet
		// TODO::: DDOS!!??
		return
	}

	// Dropping packets is preferable to waiting for packets delayed due to retransmissions.
	// Developer can ask to complete data for offline usage after first data usage.
	return
}
