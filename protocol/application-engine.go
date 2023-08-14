/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// ApplicationEngine is the interface that return some useful data about the engine that implement Application protocol
// In many ways it is like window.navigator in web ecosystem
type ApplicationEngine interface {
	Name() string
	CharacterSet() string
}
