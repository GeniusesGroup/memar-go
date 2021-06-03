/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"bytes"

	etime "../earth-time"
	er "../error"
	"../ganjine"
	"../giti"
	lang "../language"
)

const (
	chaparConnectionStructureID uint64 = 11984659004635277
)

var chaparConnectionStructure = ganjine.DataStructure{
	URN:                "urn:giti:chapar.giti:data-structure:chapar-connection",
	ID:                 11984659004635277,
	IssueDate:          1622176092,
	ExpiryDate:         0,
	ExpireInFavorOfURN: "", // Other structure name
	ExpireInFavorOfID:  0,  // Other StructureID! Handy ID or Hash of ExpireInFavorOf!
	Status:             ganjine.DataStructureStatePreAlpha,
	Structure:          Connection{},

	Detail: map[lang.Language]ganjine.DataStructureDetail{
		lang.LanguageEnglish: {
			Name:        "Chapar Connection",
			Description: ``,
			TAGS:        []string{},
		},
	},
}

// Connection ---Read locale description in chaparConnectionStructure---
type Connection struct {
	/* Common header data */
	recordID          [32]byte
	recordStructureID uint64
	recordSize        uint64
	writeTime         etime.Time
	ownerAppID        [32]byte

	/* Connection data */
	state      giti.ConnectionState
	weight     giti.ConnectionWeight
	port       giti.LinkPort `syllab:"-"`
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

// ReadFrom set Path, ReversePath and set MTU by calculate it!
func (c *Connection) ReadFrom(frame []byte) {
	var pathFromPeer Path
	pathFromPeer.ReadFrom(frame)
	c.setPath(pathFromPeer)

	// TODO::: Get ThingID from peer??
}

// setPath set Path, ReversePath and set MTU by calculate it!
func (c *Connection) setPath(pathFromPeer Path) {
	c.pathFromPeer = pathFromPeer
	c.pathToPeer = pathFromPeer.GetReverse()
	c.mtu = MaxFrameLen - int(FixedHeaderLength+pathFromPeer.LenAsByte())
}

// setAlternativePath register connection new path in the connection alternativePaths!
func (c *Connection) setAlternativePath(alternativePath Path) (err *er.Error) {
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
func (c *Connection) ChangePath(pathFromPeer Path) (err *er.Error) {
	if bytes.Equal(c.pathFromPeer, pathFromPeer) {
		err = ErrPathAlreadyUse
		return
	}
	c.setAlternativePath(c.pathFromPeer)
	c.setPath(pathFromPeer)
	return
}

// MTU return max payload size that this connection can carry on active path!
func (c *Connection) MTU() int {
	return c.mtu
}

// Send use for
func (c *Connection) Send(nexHeaderID giti.LinkHeaderID, payload giti.WriterTo) (err *er.Error) {
	var payloadLen int = payload.Len()
	if payloadLen > c.mtu {
		return ErrMTU
	}

	// TODO::: need to check path exist here to use c.AlternativePath?

	var pathLen byte = c.pathToPeer.LenAsByte()
	if pathLen == 0 {
		pathLen = MaxHopCount // broadcast frame
	}
	var payloadLoc int = 3 + int(pathLen)
	var frameLength int = payloadLoc + payloadLen
	var frame = make([]byte, frameLength)

	SetHopCount(frame, pathLen)
	SetNextHeader(frame, byte(nexHeaderID))
	c.pathToPeer.WriteTo(frame)
	err = payload.WriteTo(frame[payloadLoc:])

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

func (c *Connection) saveConn() {
	// save to PersiaOS storage in caller domain scope!
	return
}

/*
	-- Get Last Methods --
*/

// GetLastByPath return last connection by given c.Path from storage if exist
func (c *Connection) GetLastByPath() (err *er.Error) {
	return
}

// GetLastByThingID return last connection by given c.ThingID from storage if exist
func (c *Connection) GetLastByThingID() (err *er.Error) {
	return
}
