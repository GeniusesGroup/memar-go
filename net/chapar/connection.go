/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"bytes"

	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
	operation_p "memar/process/operation/protocol"
	// string_p "memar/codec/string/protocol"
)

// Connection keep some data and provide some methods to use as memar/protocol.NetworkLink
type Connection struct {
	/* Connection data */
	state      net_p.SocketStatus
	weight     operation_p.Weight
	port       *port `syllab:"-"`
	pathToPeer Path

	/* Peer data */
	pathFromPeer     Path // Chapar switch spec
	alternativePaths []Path
}

// Init set some data from given frame as connection initialize.
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (c *Connection) Init(frame Frame, port *port) (err error_p.Error) {
	err = c.pathFromPeer.Unmarshal(frame)
	if err != nil {
		return
	}
	c.pathFromPeer.CopyReverseTo(&c.pathToPeer)
	c.port = port

	return
}
func (c *Connection) Reinit() (err error_p.Error) {
	return
}
func (c *Connection) Deinit() (err error_p.Error) {
	return
}

//memar:impl memar/protocol.NetworkLink
func (c *Connection) FrameType() net_p.FrameType { return net_p.FrameType_Chapar }

//memar:impl memar/net/protocol.Field_NetworkAddresses
func (c *Connection) LocalAddr() net_p.NetworkAddress  { return &c.pathFromPeer }
func (c *Connection) RemoteAddr() net_p.NetworkAddress { return &c.pathToPeer }

func (c *Connection) ActivePaths() Path        { return c.pathToPeer }
func (c *Connection) AlternativePaths() []Path { return c.alternativePaths }

//memar:impl memar/protocol.NetworkLink
func (c *Connection) WriteFrame(packet net_p.Packet) (n int, err error_p.Error) {
	var f = Frame(packet)
	f.Init(c.pathToPeer.path[:])
	n = int(f.FrameLength())
	return
}

// Send use to send complete frame that get from c.NewFrame
//
//memar:impl memar/protocol.NetworkLink
func (c *Connection) Send(frame []byte) (err error_p.Error) {
	// send frame by connection port
	err = c.port.Send(frame)
	if err != nil {
		return
	}
	// TODO::: need to check path exist here to use c.AlternativePath?

	// c.Metric.PacketSent(uint64(len(frame)))
	return
}

func (c *Connection) ReSend(frame []byte) (err error_p.Error) {
	// send frame by connection port
	err = c.port.Send(frame)
	if err != nil {
		return
	}
	// TODO::: need to check path exist here to use c.AlternativePath?

	// c.Metric.PacketResend(uint64(len(frame)))
	return
}

// setAlternativePath register connection new path in the connection alternativePaths.
func (c *Connection) setAlternativePath(alternativePath Path) (err error_p.Error) {
	if bytes.Equal(c.pathToPeer.path[:], alternativePath.path[:]) {
		// err = &chapar_errs.PathAlreadyUse
		return
	}
	for _, path := range c.alternativePaths {
		if bytes.Equal(path.path[:], alternativePath.path[:]) {
			// err = &chapar_errs.PathAlreadyExist
			return
		}
	}
	c.alternativePaths = append(c.alternativePaths, alternativePath)
	return
}

func (c *Connection) newConnection(port *port, frame []byte) {
	c.Init(frame, port)

	return
}

func (c *Connection) establishByPath(path []byte) (err error_p.Error) {
	return
}

// ChangePath change the main connection path from alternative paths.
func (c *Connection) changePath(alternativeIndex int) (err error_p.Error) {
	var temp Path
	temp = c.pathToPeer

	c.pathToPeer = c.alternativePaths[alternativeIndex]
	c.pathFromPeer.CopyReverseTo(&c.pathToPeer)

	c.alternativePaths[alternativeIndex] = temp
	return
}
