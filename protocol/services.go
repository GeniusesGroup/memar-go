/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Services use to register services to get them in a desire way e.g. sid in http query.
type Services interface {
	// Register use to register application services.
	// Due to minimize performance impact, This method isn't safe to use concurrently and
	// must register all service before use GetService methods.
	Register(s Service) (err Error)
	Delete(s Service) (err Error)

	Services() []Service
	GetByID(sID ServiceID) (ser Service, err Error)
	GetByMediaType(mt string) (ser Service, err Error)
}
