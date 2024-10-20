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
func GetByID(sID service_p.ServiceID) (ser protocol.Service, err protocol.Error) {
	return services.GetByID(sID)
}
func GetByMediaType(mt string) (ser protocol.Service, err protocol.Error) {
	return services.GetByMediaType(mt)
}

// TODO::: decide about poolSize by hardware
const poolSizes = 512

var services = services_{
	poolByRegisterTime: make([]protocol.Service, poolSizes),
	poolByID:           make(map[service_p.ServiceID]protocol.Service, poolSizes),
	poolByMediaType:    make(map[string]protocol.Service, poolSizes),
}

type services_ struct {
	poolByRegisterTime []protocol.Service
	poolByID           map[service_p.ServiceID]protocol.Service
	poolByMediaType    map[string]protocol.Service
}

// RegisterService use to register application services.
// Due to minimize performance impact, This method isn't safe to use concurrently and
// must register all service before use GetService methods.
//
//memar:impl memar/protocol.Services
func (ss *services_) Register(s protocol.Service) (err protocol.Error) {
	if s.ServiceID() == 0 {
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
func (ss *services_) GetByID(sID service_p.ServiceID) (ser protocol.Service, err protocol.Error) {
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
	delete(ss.poolByID, s.ServiceID())
	delete(ss.poolByMediaType, s.MediaType())
	// TODO::: delete from ss.poolByRegisterTime
	return
}

func (ss *services_) registerServiceByMediaType(s protocol.Service) (err protocol.Error) {
	var serviceID = s.ServiceID()
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
