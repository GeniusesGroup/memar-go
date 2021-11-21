/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"bytes"

	etime "../earth-time"
	"../protocol"
)

// Connection ---Read locale description in chaparConnectionStructure---
type Connection struct {
	writeTime etime.Time

	/* Connection data */
	state      protocol.ConnectionState
	weight     protocol.ConnectionWeight
	port       *port `syllab:"-"`
	mtu        int
	pathToPeer Path

	/* Peer data */
	pathFromPeer     Path // Chapar switch spec
	alternativePaths []Path
	thingID          [32]byte

	/* Metrics data */
	lastUsage                  etime.Time // Last use of this connection
	bytesSent                  uint64     // Counts the bytes of frames payload sent.
	framesSent                 uint64     // Counts sent frames.
	bytesReceived              uint64     // Counts the bytes of frames payloads received.
	framesReceived             uint64     // Counts received frames.
	failedFramesSent           uint64     // Counts failed frames receive for firewalling server from some attack types!
	failedFramesReceived       uint64     // Counts failed frames receive for firewalling server from some attack types!
	notRequestedFramesReceived uint64     // Counts not requested frames received for firewalling server from some attack types!
}

// init set Path, ReversePath and set MTU by calculate it!
func (c *Connection) init(frame []byte) (err protocol.Error) {
	var pathFromPeer Path
	pathFromPeer.CopyFrom(frame)
	c.setPath(pathFromPeer)

	// TODO::: Get ThingID from peer??
}

// MTU return max payload size that this connection can carry on active path!
func (c *Connection) MTU() int {
	return c.mtu
}

// Send use for
func (c *Connection) Send(nexHeaderID protocol.NetworkLinkNextHeaderID, payload protocol.Codec) (err protocol.Error) {
	var payloadLen int = payload.Len()

	// TODO::: need to check path exist here to use c.AlternativePath?

	frame = c.newFrame(nexHeaderID, payload, payloadLen)

	// send frame by connection port!
	err = c.port.Send(frame)

	// Add metrics data
	c.lastUsage = etime.Now()
	if err != nil {
		c.failedFramesSent++
	} else {
		c.bytesSent += uint64(payloadLen)
		c.framesSent++
	}
	return
}

// SendAsync use to send the frame async!
func (c *Connection) SendAsync(nexHeaderID protocol.NetworkLinkNextHeaderID, payload protocol.Codec) (err protocol.Error) {
	var payloadLen int = payload.Len()

	// TODO::: need to check path exist here to use c.AlternativePath?

	frame = c.newFrame(nexHeaderID, payload, payloadLen)

	// send frame by connection port!
	err = c.port.SendAsync(frame)

	// Add metrics data
	c.lastUsage = etime.Now()
	c.bytesSent += uint64(payloadLen)
	c.framesSent++
	return
}

// newFrame makes new unicast||broadcast frame!
func (c *Connection) newFrame(nexHeaderID protocol.NetworkLinkNextHeaderID, payload protocol.Codec, payloadLen int) (frame []byte) {
	if payloadLen > c.mtu {
		return ErrMTU
	}

	var pathLen byte = c.pathToPeer.LenAsByte()
	var payloadLoc int = 3 + int(pathLen)
	var frameLength int = payloadLoc + payloadLen
	frame = make([]byte, frameLength)

	SetHopCount(frame, pathLen)
	SetNextHeader(frame, byte(nexHeaderID))
	c.pathToPeer.MarshalTo(frame)
	payload.MarshalTo(frame[payloadLoc:])
	return
}

// setPath set Path, ReversePath and set MTU by calculate it!
func (c *Connection) setPath(pathFromPeer Path) {
	c.pathFromPeer = pathFromPeer
	c.pathToPeer = pathFromPeer.GetReverse()
	c.mtu = MaxFrameLen - int(FixedHeaderLength+pathFromPeer.LenAsByte())
}

// setAlternativePath register connection new path in the connection alternativePaths!
func (c *Connection) setAlternativePath(alternativePath Path) (err protocol.Error) {
	for path := range c.alternativePaths {
		if bytes.Equal(path, alternativePath) {
			err = ErrPathAlreadyExist
			return
		}
	}
	c.alternativePaths = append(c.alternativePaths, alternativePath)
	return
}

// ChangePath change the main connection path!
func (c *Connection) changePath(pathFromPeer Path) (err protocol.Error) {
	if bytes.Equal(c.pathFromPeer, pathFromPeer) {
		err = ErrPathAlreadyUse
		return
	}
	c.setAlternativePath(c.pathFromPeer)
	c.setPath(pathFromPeer)
	return
}
