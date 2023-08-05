/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"bytes"

	"memar/protocol"
)

// Connection keep some data and provide some methods to use as memar/protocol.NetworkLink
type Connection struct {
	/* Connection data */
	state      protocol.NetworkStatus
	weight     protocol.Weight
	port       *port `syllab:"-"`
	pathToPeer Path

	/* Peer data */
	pathFromPeer     Path // Chapar switch spec
	alternativePaths []Path
	thingID          protocol.UserUUID
}

// Init set some data from given frame as connection initialize.
//
//memar:impl memar/protocol.ObjectLifeCycle
func (c *Connection) Init(frame Frame, port *port) (err protocol.Error) {
	err = c.pathFromPeer.Unmarshal(frame)
	if err != nil {
		return
	}
	c.pathFromPeer.CopyReverseTo(&c.pathToPeer)
	c.port = port

	// TODO::: Get ThingID from peer??

	return
}
func (c *Connection) Reinit() (err protocol.Error) {
	return
}
func (c *Connection) Deinit() (err protocol.Error) {
	return
}

//memar:impl memar/protocol.NetworkLink
func (c *Connection) FrameType() protocol.Network_FrameType { return protocol.Network_FrameType_Chapar }

//memar:impl memar/protocol.NetworkAddress
func (c *Connection) LocalAddr() protocol.Stringer  { return &c.pathFromPeer }
func (c *Connection) RemoteAddr() protocol.Stringer { return &c.pathToPeer }

func (c *Connection) ActivePaths() Path        { return c.pathToPeer }
func (c *Connection) AlternativePaths() []Path { return c.alternativePaths }

//memar:impl memar/protocol.NetworkLink
func (c *Connection) WriteFrame(packet []byte) (n int, err protocol.Error) {
	var f = Frame(packet)
	f.Init(c.pathToPeer.path[:])
	n = f.FrameLen()
	return
}

// Send use to send complete frame that get from c.NewFrame
//
//memar:impl memar/protocol.NetworkLink
func (c *Connection) Send(frame []byte) (err protocol.Error) {
	// send frame by connection port
	err = c.port.Send(frame)
	if err != nil {
		return
	}
	// TODO::: need to check path exist here to use c.AlternativePath?

	// c.Metric.PacketSent(uint64(len(frame)))
	return
}

func (c *Connection) ReSend(frame []byte) (err protocol.Error) {
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
func (c *Connection) setAlternativePath(alternativePath Path) (err protocol.Error) {
	if bytes.Equal(c.pathToPeer.path[:], alternativePath.path[:]) {
		err = &ErrPathAlreadyUse
		return
	}
	for _, path := range c.alternativePaths {
		if bytes.Equal(path.path[:], alternativePath.path[:]) {
			err = &ErrPathAlreadyExist
			return
		}
	}
	c.alternativePaths = append(c.alternativePaths, alternativePath)
	return
}

func (c *Connection) newConnection(port *port, frame []byte) {
	c.Init(frame, port)

	// TODO::: get ThingID from peer or func args??

	return
}

func (c *Connection) establishByPath(path []byte) (err protocol.Error) {
	return
}

func (c *Connection) establishByThingID(thingID [32]byte) (err protocol.Error) {
	return
}

// ChangePath change the main connection path from alternative paths.
func (c *Connection) changePath(alternativeIndex int) (err protocol.Error) {
	var temp Path
	temp = c.pathToPeer

	c.pathToPeer = c.alternativePaths[alternativeIndex]
	c.pathFromPeer.CopyReverseTo(&c.pathToPeer)

	c.alternativePaths[alternativeIndex] = temp
	return
}
