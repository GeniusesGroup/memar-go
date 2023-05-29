/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Modules is the interface that must implement by any Application.
type Modules interface {
	// RegisterModule use to register application modules.
	// Due to minimize performance impact, This method isn't safe to use concurrently and
	// must register all module before use GetModule methods.
	RegisterModule(m Module)

	Modules() []Module
	GetModuleByID(id MediaTypeID) (m Module, err Error)
	GetModuleByMediaType(mt string) (m Module, err Error)
}

// Module is the interface that must implement by any struct to be a module.
type Module interface {
	Source() string // Domain, URL, ...
	// Two type of depend on: use other module or wrapper for some modules to introduce as one package.
	DependOn() []Module

	Services() []Service
	Pages() []GUIPage
	// Widgets() []GUI

	Details
	MediaType
	ObjectLifeCycle
}
