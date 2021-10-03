/* For license and copyright information please see LEGAL file in repository */

package object

import (
	"../protocol"
)

const (
	DomainName = "object.protocol"
)

// registerServices use to initialize services and register them in the application.
func registerServices() {
	protocol.App.RegisterService(&DeleteService)
	protocol.App.RegisterService(&GetMetadataService)
	protocol.App.RegisterService(&GetService)
	protocol.App.RegisterService(&ReadService)
	protocol.App.RegisterService(&SaveService)
	protocol.App.RegisterService(&WipeService)
	protocol.App.RegisterService(&WriteService)
}

// RequestType indicate request type!
type RequestType uint8

// Services request types
const (
	RequestTypeStandalone RequestType = iota
	RequestTypeBroadcast
)
