/* For license and copyright information please see the LEGAL file in the code repository */

package services_p

import (
	service_p "memar/operation/service/protocol"
	"memar/protocol"
)

// Services use to register services to get them in a desire way e.g. sid in http query.
type Services interface {
	// Register use to register application services.
	// Due to minimize performance impact, This method isn't safe to use concurrently and
	// must register all service before use GetService methods.
	Register(s service_p.Service) (err protocol.Error)
	Delete(s service_p.Service) (err protocol.Error)

	Services() []service_p.Service
	GetByID(sID service_p.ServiceID) (ser service_p.Service, err protocol.Error)
	GetByMediaType(mt string) (ser service_p.Service, err protocol.Error)
}
