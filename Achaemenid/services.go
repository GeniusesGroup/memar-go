/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "../log"

type services struct {
	idPool  map[uint32]*Service
	uriPool map[string]*Service
}

func (ss *services) init() {
	if ss.idPool == nil {
		ss.idPool = make(map[uint32]*Service)
	}
	if ss.uriPool == nil {
		ss.uriPool = make(map[string]*Service)
	}
}

// RegisterService use to set or change specific service detail!
func (ss *services) RegisterService(st *Service) {
	if st.ID == 0 {
		log.Warn("Service '", st.Name, "', give 0 as service ID! it won't register to use by ID! legal ID must hash of service name")
	} else {
		var es = ss.GetServiceHandlerByID(st.ID)
		if es != nil {
			// Warn developer this Service use ID for other service and panic app
			log.Warn(st.Name+" service with ", st.ID, " ID, Used before for other service and it illegal to reuse IDs")
			log.Fatal("Exiting service name is " + es.Name)
		} else {
			ss.idPool[st.ID] = st
		}
	}

	if len(st.URI) != 0 {
		if ss.GetServiceHandlerByURI(st.URI) != nil {
			// Warn developer this Service use ID for other service and this panic
			log.Fatal(st.Name+" service with ", st.URI, " URI, Used before for other service and it illegal to reuse URI")
		} else {
			ss.uriPool[st.URI] = st
		}
	}
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
