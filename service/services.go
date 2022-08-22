/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

// Services store all application service
type Services struct {
	poolByID        map[protocol.MediaTypeID]protocol.Service
	poolByMediaType map[string]protocol.Service
	poolByURIPath   map[string]protocol.Service
}

// Init use to initialize
func (ss *Services) Init() {
	const poolSizes = 512
	// TODO::: decide about poolSize by hardware

	ss.poolByID = make(map[protocol.MediaTypeID]protocol.Service, poolSizes)
	ss.poolByURIPath = make(map[string]protocol.Service, poolSizes)
	ss.poolByMediaType = make(map[string]protocol.Service, poolSizes)
}

// RegisterService use to register application services.
// Due to minimize performance impact, This method isn't safe to use concurrently and
// must register all service before use GetService methods.
func (ss *Services) RegisterService(s protocol.Service) {
	if s.ID() == 0 && s.URI() == "" {
		// This condition must be true just in the dev phase.
		panic("Service must have a valid URI or mediatype. It is rule to add more detail about service. Initialize inner s.MediaType.Init() first if use libgo/service package")
	}

	ss.registerServiceByMediaType(s)
	ss.registerServiceByURI(s)
}

func (ss *Services) registerServiceByMediaType(s protocol.Service) {
	var serviceID = s.ID()
	var exitingServiceByID = ss.GetServiceByID(serviceID)
	if exitingServiceByID != nil {
		// This condition will just be true in the dev phase.
		panic("ID associated for '" + s.MediaType() + "' Used before for other service and not legal to reuse same ID for other services\n" +
			"Exiting service MediaType is: " + exitingServiceByID.MediaType())
	} else {
		ss.poolByID[serviceID] = s
	}

	var serviceMediaType = s.MediaType()
	var exitingServiceByMediaType = ss.GetServiceByMediaType(serviceMediaType)
	if exitingServiceByMediaType != nil {
		// This condition will just be true in the dev phase.
		panic("This mediatype '" + serviceMediaType + "' register already before for other service and not legal to reuse same mediatype for other services\n")
	} else {
		ss.poolByMediaType[serviceMediaType] = s
	}
}

func (ss *Services) registerServiceByURI(s protocol.Service) {
	var serviceURI = s.URI()
	if serviceURI != "" {
		var exitingServiceByURI = ss.GetServiceByURI(serviceURI)
		if exitingServiceByURI != nil {
			// This condition will just be true in the dev phase.
			panic("URI associated for '" + s.MediaType() + " service with `" + serviceURI + "` as URI, Used before for other service and not legal to reuse URI for other services\n" +
				"Exiting service MediaType is: " + exitingServiceByURI.MediaType())
		} else {
			ss.poolByMediaType[serviceURI] = s
		}
	}
}

// GetServiceByID use to get specific service handler by service ID
func (ss *Services) GetServiceByID(serviceID protocol.MediaTypeID) protocol.Service {
	return ss.poolByID[serviceID]
}

// GetServiceByMediaType use to get specific service handler by service URI
func (ss *Services) GetServiceByMediaType(mt string) protocol.Service {
	return ss.poolByMediaType[mt]
}

// GetServiceByURI use to get specific service handler by service URI path
func (ss *Services) GetServiceByURI(uri string) protocol.Service {
	return ss.poolByURIPath[uri]
}

// DeleteService use to delete specific service in services list.
func (ss *Services) DeleteService(s protocol.Service) {
	delete(ss.poolByID, s.ID())
	delete(ss.poolByMediaType, s.MediaType())
	delete(ss.poolByMediaType, s.URI())
}
