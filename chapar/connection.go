/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"bytes"

	"libgo/connection"
	"libgo/protocol"
)

// Connection keep some data and provide some methods to use as /libgo/protocol.NetworkLink_Connection
type Connection struct {
	/* Connection data */
	state      protocol.NetworkStatus
	weight     protocol.Weight
	port       *port `syllab:"-"`
	mtu        int   // max payload size that this connection can carry on active path!
	pathToPeer Path

	/* Peer data */
	pathFromPeer     Path // Chapar switch spec
	alternativePaths []Path
	thingID          [32]byte

	connection.Metric
}

// Init set some data from given frame as connection initialize.
func (c *Connection) Init(frame Frame, port *port) (err protocol.Error) {
	err = c.pathFromPeer.Unmarshal(frame)
	if err != nil {
		return
	}
	c.pathFromPeer.CopyReverseTo(&c.pathToPeer)
	c.setMTU()
	c.port = port

	// TODO::: Get ThingID from peer??

	return
}

func (c *Connection) MTU() int                 { return c.mtu }
func (c *Connection) ActivePaths() Path        { return c.pathToPeer }
func (c *Connection) AlternativePaths() []Path { return c.alternativePaths }

// NewFrame makes new unicast||broadcast frame.
func (c *Connection) NewFrame(nexHeaderID protocol.NetworkLink_NextHeaderID, payloadLen int) (frame []byte, payload []byte, err protocol.Error) {
	if payloadLen > c.mtu {
		err = &ErrMTU
		return
	}

	var nhID = NetworkLink_NextHeaderIDToChaparNextHeaderID(nexHeaderID)

	var f Frame
	f.Init(nhID, c.pathToPeer.path[:], payloadLen)

	payload = f.Payload()[:0]
	frame = f
	return
}

// Send use to send complete frame that get from c.NewFrame
func (c *Connection) Send(frame []byte) (err protocol.Error) {
	// send frame by connection port
	err = c.port.Send(frame)
	if err != nil {
		return
	}
	// TODO::: need to check path exist here to use c.AlternativePath?

	c.Metric.PacketSent(uint64(len(frame)))
	return
}

func (c *Connection) ReSend(frame []byte) (err protocol.Error) {
	// send frame by connection port
	err = c.port.Send(frame)
	if err != nil {
		return
	}
	// TODO::: need to check path exist here to use c.AlternativePath?

	c.Metric.PacketResend(uint64(len(frame)))
	return
}

// setMTU set MTU by calculate it from path length.
func (c *Connection) setMTU() {
	c.mtu = MaxFrameLen - int(fixedHeaderLength+c.pathToPeer.LenAsByte())
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
	c.setMTU()

	c.alternativePaths[alternativeIndex] = temp
	return
}
