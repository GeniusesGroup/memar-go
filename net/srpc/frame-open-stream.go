/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"../protocol"
	"../syllab"
)

/*
type openStreamFrame struct {
	ProtocolID  uint16 // protocol ID usage is like TCP||UDP ports that indicate payload protocol.
	SerErrID    uint64 // Service or Error
	CompressID  uint64
	TotalPacket uint32 // Expected packets count that send over this stream!
	DataLength  uint64
	Weight      protocol.ConnectionWeight
}
*/
type openStreamFrame []byte

func (f openStreamFrame) ProtocolID() uint16                { return syllab.GetUInt16(f, 0) }
func (f openStreamFrame) SerErrID() uint64                  { return syllab.GetUInt64(f, 2) }
func (f openStreamFrame) CompressID() uint64                { return syllab.GetUInt64(f, 10) }
func (f openStreamFrame) Weight() protocol.ConnectionWeight { return protocol.ConnectionWeight(f[10]) }
func (f openStreamFrame) NextFrame() []byte                 { return f[18:] }

// setStreamSettings set stream settings like time sensitive use in VoIP, IPTV, ...
func openStream(conn protocol.Connection, frame openStreamFrame) (err protocol.Error) {
	// TODO::: allow multiple settings set??

	// Check server supported requested protocol
	var ProtocolID = protocol.NetworkApplicationProtocolID(frame.ProtocolID())
	var protocolHandler protocol.NetworkApplicationHandler = protocol.App.GetNetworkApplicationHandler(ProtocolID)
	if protocolHandler == nil {
		// Send response or just ignore packet
		// TODO::: DDOS!!??
		return
	}

	// Dropping packets is preferable to waiting for packets delayed due to retransmissions.
	// Developer can ask to complete data for offline usage after first data usage.
	return
}
