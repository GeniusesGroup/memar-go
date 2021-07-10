/* For license and copyright information please see LEGAL file in repository */

package giti

// Stream is the interface that must implement by any struct to be Stream!
type Streams interface {
	GetStreamByID(ID uint32) Stream
	RegisterStream(Stream)
	CloseStream(Stream)
}
