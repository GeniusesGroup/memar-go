/* For license and copyright information please see the LEGAL file in the code repository */

package services

import (
	"memar/audit/log"
	error_p "memar/error/protocol"
	service_p "memar/operation/service/protocol"
	errs "memar/operation/services/errors"
)

func Register(s service_p.Service) (err error_p.Error) { return services.Register(s) }
func Delete(s service_p.Service) (err error_p.Error)   { return services.Delete(s) }
func Services() []service_p.Service                    { return services.Services() }
func GetByID(sID service_p.ID) (ser service_p.Service, err error_p.Error) {
	return services.GetByID(sID)
}
func GetByMediaType(mt string) (ser service_p.Service, err error_p.Error) {
	return services.GetByMediaType(mt)
}

// TODO::: decide about poolSize by hardware
const poolSizes = 512

var services = services_{
	poolByRegisterTime: make([]service_p.Service, poolSizes),
	poolByID:           make(map[service_p.ID]service_p.Service, poolSizes),
	poolByMediaType:    make(map[string]service_p.Service, poolSizes),
}

type services_ struct {
	poolByRegisterTime []service_p.Service
	poolByID           map[service_p.ID]service_p.Service
	poolByMediaType    map[string]service_p.Service
}

// RegisterService use to register application services.
// Due to minimize performance impact, This method isn't safe to use concurrently and
// must register all service before use GetService methods.
//
//memar:impl memar/protocol.Services
func (self *services_) Register(s service_p.Service) (err error_p.Error) {
	if s.ServiceID() == 0 {
		err = &errs.ServiceNotProvideIdentifier
		return
	}

	self.registerServiceByMediaType(s)
	self.poolByRegisterTime = append(self.poolByRegisterTime, s)
	return
}

// Services use to get all services registered.
//
//memar:impl memar/protocol.Services
func (self *services_) Services() []service_p.Service { return self.poolByRegisterTime }

// GetServiceByID use to get specific service handler by service ID
//
//memar:impl memar/protocol.Services
func (self *services_) GetByID(sID service_p.ID) (ser service_p.Service, err error_p.Error) {
	ser = self.poolByID[sID]
	if ser == nil {
		err = &errs.NotFound
	}
	return
}

// GetServiceByMediaType use to get specific service handler by service URI
//
//memar:impl memar/protocol.Services
func (self *services_) GetByMediaType(mt string) (ser service_p.Service, err error_p.Error) {
	ser = self.poolByMediaType[mt]
	if ser == nil {
		err = &errs.NotFound
	}
	return
}

// DeleteService use to delete specific service in services list.
func (self *services_) Delete(s service_p.Service) (err error_p.Error) {
	delete(self.poolByID, s.ServiceID())
	delete(self.poolByMediaType, s.MediaType())
	// TODO::: delete from self.poolByRegisterTime
	return
}

func (self *services_) registerServiceByMediaType(s service_p.Service) (err error_p.Error) {
	var serviceID = s.ServiceID()
	var exitingServiceByID, _ = self.GetByID(serviceID)
	if exitingServiceByID != nil {
		err = &errs.ServiceDuplicateIdentifier
		log.Fatal(s, "ID associated for '"+s.MediaType()+"' Used before for other service and not legal to reuse same ID for other services\n"+
			"	Exiting service MediaType is: "+exitingServiceByID.MediaType())
	} else {
		self.poolByID[serviceID] = s
	}

	var serviceMediaType = s.MediaType()
	var exitingServiceByMediaType, _ = self.GetByMediaType(serviceMediaType)
	if exitingServiceByMediaType != nil {
		err = &errs.ServiceDuplicateIdentifier
		log.Fatal(s, "This mediatype '"+serviceMediaType+"' register already before for other service and not legal to reuse same mediatype for other services\n")
	} else {
		self.poolByMediaType[serviceMediaType] = s
	}
	return
}
