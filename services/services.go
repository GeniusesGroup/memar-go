/* For license and copyright information please see the LEGAL file in the code repository */

package services

import (
	"memar/log"
	"memar/protocol"
	errs "memar/services/errors"
)

func Register(s protocol.Service) (err protocol.Error) { return services.Register(s) }
func Delete(s protocol.Service) (err protocol.Error)   { return services.Delete(s) }
func Services() []protocol.Service                     { return services.Services() }
func GetByID(sID protocol.ServiceID) (ser protocol.Service, err protocol.Error) {
	return services.GetByID(sID)
}
func GetByMediaType(mt string) (ser protocol.Service, err protocol.Error) {
	return services.GetByMediaType(mt)
}

var services services_

type services_ struct {
	poolByRegisterTime []protocol.Service
	poolByID           map[protocol.ServiceID]protocol.Service
	poolByMediaType    map[string]protocol.Service
}

// Init use to initialize
func (ss *services_) Init() (err protocol.Error) {
	const poolSizes = 512
	// TODO::: decide about poolSize by hardware

	ss.poolByRegisterTime = make([]protocol.Service, poolSizes)
	ss.poolByID = make(map[protocol.ServiceID]protocol.Service, poolSizes)
	ss.poolByMediaType = make(map[string]protocol.Service, poolSizes)
	return
}
func (ss *services_) Reinit() (err protocol.Error) {
	// TODO:::
	// for _, s := range ss.poolByURIPath {
	// 	s.Reinit()
	// }
	return
}
func (ss *services_) Deinit() (err protocol.Error) {
	for _, s := range ss.poolByRegisterTime {
		err = s.Deinit()
		// TODO::: easily return if occur any error??
		if err != nil {
			return
		}
	}
	return
}

// RegisterService use to register application services.
// Due to minimize performance impact, This method isn't safe to use concurrently and
// must register all service before use GetService methods.
//
//memar:impl memar/protocol.Services
func (ss *services_) Register(s protocol.Service) (err protocol.Error) {
	if s.ID() == 0 {
		err = &errs.ErrServiceNotProvideIdentifier
		return
	}

	ss.registerServiceByMediaType(s)
	ss.poolByRegisterTime = append(ss.poolByRegisterTime, s)
	return
}

// Services use to get all services registered.
//
//memar:impl memar/protocol.Services
func (ss *services_) Services() []protocol.Service { return ss.poolByRegisterTime }

// GetServiceByID use to get specific service handler by service ID
//
//memar:impl memar/protocol.Services
func (ss *services_) GetByID(sID protocol.ServiceID) (ser protocol.Service, err protocol.Error) {
	ser = ss.poolByID[sID]
	if ser == nil {
		err = &errs.ErrNotFound
	}
	return
}

// GetServiceByMediaType use to get specific service handler by service URI
//
//memar:impl memar/protocol.Services
func (ss *services_) GetByMediaType(mt string) (ser protocol.Service, err protocol.Error) {
	ser = ss.poolByMediaType[mt]
	if ser == nil {
		err = &errs.ErrNotFound
	}
	return
}

// DeleteService use to delete specific service in services list.
func (ss *services_) Delete(s protocol.Service) (err protocol.Error) {
	delete(ss.poolByID, s.ID())
	delete(ss.poolByMediaType, s.MediaType())
	// TODO::: delete from ss.poolByRegisterTime
	return
}

func (ss *services_) registerServiceByMediaType(s protocol.Service) (err protocol.Error) {
	var serviceID = s.ID()
	var exitingServiceByID, _ = ss.GetByID(serviceID)
	if exitingServiceByID != nil {
		err = &errs.ErrServiceDuplicateIdentifier
		log.Fatal(s, "ID associated for '"+s.MediaType()+"' Used before for other service and not legal to reuse same ID for other services\n"+
			"	Exiting service MediaType is: "+exitingServiceByID.MediaType())
	} else {
		ss.poolByID[serviceID] = s
	}

	var serviceMediaType = s.MediaType()
	var exitingServiceByMediaType, _ = ss.GetByMediaType(serviceMediaType)
	if exitingServiceByMediaType != nil {
		err = &errs.ErrServiceDuplicateIdentifier
		log.Fatal(s, "This mediatype '"+serviceMediaType+"' register already before for other service and not legal to reuse same mediatype for other services\n")
	} else {
		ss.poolByMediaType[serviceMediaType] = s
	}
	return
}
