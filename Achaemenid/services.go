/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// Services store server services!
type Services struct {
	Services []*Service // Use RegisterService func to register service here not directly append them!
	ByID     map[uint32]StreamHandler
	ByURI    map[string]StreamHandler
}

func (ss *Services) init() {
	if ss.ByID == nil {
		ss.ByID = make(map[uint32]StreamHandler)
	}
	if ss.ByURI == nil {
		ss.ByURI = make(map[string]StreamHandler)
	}
}

// RegisterService use to set or change specific service detail!
func (ss *Services) RegisterService(st *Service) {
	if st.ID == 0 {
		Log("This service: ", st.Name, ", give 0 as service ID! it illegal to use 0 as ID! It must hash of service name")
		panic("ChaparKhane occur panic situation due to ^")
	}

	_, ok := ss.ByID[st.ID]
	if ok {
		// Warn developer this ServiceID use for other service and this panic
		Log("This ID: ", st.ID, ", Used before for other service and it illegal to reuse IDs")
		panic("ChaparKhane occur panic situation due to ^")
	} else {
		ss.ByID[st.ID] = st.Handler
		ss.Services = append(ss.Services, st)
	}
}

// GetServiceHandlerByID use to get specific service handler by service ID!
func (ss *Services) GetServiceHandlerByID(serviceID uint32) (StreamHandler, bool) {
	h, ok := ss.ByID[serviceID]
	return h, ok
}

// GetServiceHandlerByURI use to get specific service handler by service ID!
func (ss *Services) GetServiceHandlerByURI(URI string) (StreamHandler, bool) {
	var h, ok = ss.ByURI[URI]
	return h, ok
}
