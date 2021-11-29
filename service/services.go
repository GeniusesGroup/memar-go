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
	var existingService = ss.GetServiceByID(serviceID)
	if existingService != nil {
		// Warn developer this Service use ID for other service and panic app
		protocol.App.Log(protocol.LogType_Warning, s.URN().URI()+" service with ", serviceID, " ID, Used before for other service and it illegal to reuse IDs")
		protocol.App.LogFatal("Exiting service URN is " + existingService.URN().URI())
	} else {
		ss.poolByID[serviceID] = s
	}
}

func (ss *Services) registerServiceByURI(s protocol.Service) {
	if len(s.URI()) != 0 {
		if ss.GetServiceByURI(s.URN().Domain(), s.URI()) != nil {
			// Warn developer this Service use ID for other service and this panic
			protocol.App.LogFatal(s.URN().URI()+" service with ", s.URI(), " URI, Used before for other service and it illegal to reuse URI")
		} else {
			ss.poolByURI[s.URI()] = s
		}
	}
}

// GetServiceByID use to get specific service handler by service ID!
func (ss *Services) GetServiceByID(serviceID uint64) protocol.Service {
	return ss.poolByID[serviceID]
}

// GetServiceByURI use to get specific service handler by service URI!
func (ss *Services) GetServiceByURI(domain, uri string) protocol.Service {
	return ss.poolByURI[uri]
}

// DeleteService use to delete specific service in services list.
func (ss *Services) DeleteService(s protocol.Service) {
	delete(ss.poolByID, s.URN().ID())
	delete(ss.poolByURI, s.URI())
}
