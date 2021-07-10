/* For license and copyright information please see LEGAL file in repository */

package giti

// Services is the interface that must implement by any Application!
type Services interface {
	RegisterService(s Service)
	DeleteService(s Service)
	GetServiceByID(id uint64) Service
	GetServiceByURI(uri string) Service
}
