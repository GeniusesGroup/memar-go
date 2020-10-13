/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"encoding/json"

	lang "../language"
	"../log"
)

type services struct {
	idPool  map[uint32]*Service
	uriPool map[string]*Service
}

func (ss *services) init() {
	if ss.idPool == nil {
		ss.idPool = make(map[uint32]*Service, 1000)
	}
	if ss.uriPool == nil {
		ss.uriPool = make(map[string]*Service, 1000)
	}
}

// RegisterService use to set or change specific service detail!
func (ss *services) RegisterService(s *Service) {
	var serviceName = s.Name[lang.EnglishLanguage]
	if serviceName == "" {
		log.Warn("Service with ID ((", s.ID, ")) don't has an english name! It is better to add more detail about service in english language.")
	}

	if s.ID == 0 {
		log.Warn("Service '", serviceName, "', give 0 as service ID! it won't register to use by ID! legal ID must hash of service name")
	} else {
		var es = ss.GetServiceHandlerByID(s.ID)
		if es != nil {
			// Warn developer this Service use ID for other service and panic app
			log.Warn(serviceName+" service with ", s.ID, " ID, Used before for other service and it illegal to reuse IDs")
			log.Fatal("Exiting service name is " + es.Name[lang.EnglishLanguage])
		} else {
			ss.idPool[s.ID] = s
		}
	}

	if len(s.URI) != 0 {
		if ss.GetServiceHandlerByURI(s.URI) != nil {
			// Warn developer this Service use ID for other service and this panic
			log.Fatal(serviceName+" service with ", s.URI, " URI, Used before for other service and it illegal to reuse URI")
		} else {
			ss.uriPool[s.URI] = s
		}
	}

	// s.Syllab, _ = syllab.Marshal(s, 4)
	s.JSON, _ = json.Marshal(s)
}

// GetServiceHandlerByID use to get specific service handler by service ID!
func (ss *services) GetServiceHandlerByID(serviceID uint32) *Service {
	return ss.idPool[serviceID]
}

// GetServiceHandlerByURI use to get specific service handler by service URI!
func (ss *services) GetServiceHandlerByURI(uri string) *Service {
	return ss.uriPool[uri]
}

// DeleteService use to delete specific service in services list.
func (ss *services) DeleteService(s *Service) {
	delete(ss.idPool, s.ID)
	delete(ss.uriPool, s.URI)
}
