/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// services store server services!
type services struct {
	services []*Service
	byID     map[uint32]PacketHandler
}

func (ss *services) init() {
	if ss.byID == nil {
		ss.byID = make(map[uint32]PacketHandler)
	}
}

// RegisterService use to set or change specific service detail!
func (ss *services) RegisterService(s *Service) {
	if s.ID == 0 {
		Log("This service: ", s.Name, ", give 0 as service ID! it illegal to use 0 as ID! It must hash of service name")
		panic("ChaparKhane occur panic situation due to ^")
	}

	_, ok := ss.byID[s.ID]
	if ok {
		// Warn developer this ServiceID use for other service and this panic
		Log("This ID: ", s.ID, ", Used before for other service and it illegal to reuse IDs")
		panic("ChaparKhane occur panic situation due to ^")
	} else {
		ss.byID[s.ID] = s.Handler
		ss.services = append(ss.services, s)
	}
}

// GetServiceHandlerByID use to get specific service handler by service ID!
func (ss *services) GetServiceHandlerByID(serviceID uint32) (PacketHandler, bool) {
	h, ok := ss.byID[serviceID]
	return h, ok
}
