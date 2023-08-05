/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	"memar/binary"
	"memar/protocol"
)

/*
	type OpenStreamFrame struct {
		HandlerID   uint64 // protocol ID usage is like TCP||UDP ports that indicate payload protocol.
		ServiceID   uint64
		CompressID  uint64
		DataLength  uint64
		Weight      protocol.Weight
	}

TotalPacket uint32 // Expected packets count that send over this stream.
*/
type OpenStreamFrame []byte

func (f OpenStreamFrame) HandlerID() uint64       { return binary.LittleEndian(f[0:]).Uint64() }
func (f OpenStreamFrame) ServiceID() uint64       { return binary.LittleEndian(f[2:]).Uint64() }
func (f OpenStreamFrame) CompressID() uint64      { return binary.LittleEndian(f[10:]).Uint64() }
func (f OpenStreamFrame) DataLength() uint64      { return binary.LittleEndian(f[18:]).Uint64() }
func (f OpenStreamFrame) Weight() protocol.Weight { return protocol.Weight(f[24]) }

//memar:impl memar/protocol.Network_Frame
func (f OpenStreamFrame) NextFrame() []byte { return f[25:] }

func (f OpenStreamFrame) Do(sk protocol.Socket) (err protocol.Error) {
	// TODO::: allow multiple settings set??

	// Check server supported requested protocol
	var serviceID = protocol.ServiceID(f.ServiceID())
	_, err = protocol.App.GetServiceByID(serviceID)
	if err != nil {
		// Send response or just ignore packet
		// TODO::: DDOS!!??
		return
	}

	// Dropping packets is preferable to waiting for packets delayed due to retransmissions.
	// Developer can ask to complete data for offline usage after first data usage.
	return
}
