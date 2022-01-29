/* For license and copyright information please see LEGAL file in repository */

package service

import (
	"../protocol"
)

// Services store all application service
type Services struct {
	poolByID  map[uint64]protocol.Service
	poolByURI map[string]protocol.Service
}

// Init use to initialize
func (ss *Services) Init() {
	ss.poolByID = make(map[uint64]protocol.Service, 512)
	ss.poolByURI = make(map[string]protocol.Service, 512)
}

// RegisterService use to set or change specific service detail!
func (ss *Services) RegisterService(s protocol.Service) {
	ss.registerServiceByURN(s)
	ss.registerServiceByURI(s)
	// s.ToSyllab()
	// s.ToJSON()
}

func (ss *Services) registerServiceByURN(s protocol.Service) {
	var serviceID = s.URN().ID()
	if protocol.AppMode_Dev && ss.GetServiceByID(serviceID) != nil {
		// This condition will just be true in the dev phase.
		panic("ID associated for '" + s.URN().URI() + "' Used before for other service and not legal to reuse same ID for other services\n" +
			"Exiting service URN is: " + ss.GetServiceByID(serviceID).URN().URI())
	} else {
		ss.poolByID[serviceID] = s
	}
}

func (ss *Services) registerServiceByURI(s protocol.Service) {
	var serviceURI = s.URI()
	if serviceURI != "" {
		if protocol.AppMode_Dev && ss.GetServiceByURI(serviceURI) != nil {
			// This condition will just be true in the dev phase.
			panic("URI associated for '" + s.URN().URI() + " service with `" + serviceURI + "` as URI, Used before for other service and not legal to reuse URI for other services\n" +
				"Exiting service URN is: " + ss.GetServiceByURI(serviceURI).URN().URI())
		} else {
			ss.poolByURI[serviceURI] = s
		}
	}
}

// GetServiceByID use to get specific service handler by service ID!
func (ss *Services) GetServiceByID(serviceID uint64) protocol.Service {
	return ss.poolByID[serviceID]
}

// GetServiceByURI use to get specific service handler by service URI!
func (ss *Services) GetServiceByURI(uri string) protocol.Service {
	return ss.poolByURI[uri]
}

// DeleteService use to delete specific service in services list.
func (ss *Services) DeleteService(s protocol.Service) {
	delete(ss.poolByID, s.URN().ID())
	delete(ss.poolByURI, s.URI())
}
