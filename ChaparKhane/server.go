/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// DefaultServer use as default server.
var DefaultServer = &defaultServer
var defaultServer Server

// Server store
type Server struct {
	Status            int // States locate in this file.
	Manifest          Manifest
	Services          services
	UserConnections   userConnections
	RouterConnections routerConnections
	Security          security
}

// Server Status
const (
	// ServerStateStop indicate server had been stopped
	ServerStateStop int = iota
	// ServerStateRunning indicate server is working
	ServerStateRunning
	// ServerStateStopping indicate server want to stop
	ServerStateStopping
	// ServerStateStarting indicate server plan to start and working on it
	ServerStateStarting
)

// Init method use to initialize related object with default data to prevent from panic!
func (s *Server) Init() {
	if s == nil {
		s = DefaultServer
	}
	s.Services.init()
	s.UserConnections.init()
}

// Start will start the server.
func (s *Server) Start() (err error) {

	return
}

// New XP node calculate path (latency and capacity) to all exiting node and send new information to all exiting node
// XPs can recalculate path and tell others about any change just if any change to their physical links!
// - multiple routes to the same place to be assigned the same cost and will cause traffic to be distributed evenly over those routes

// var err error
// err = gp.CheckPacket()
// if err != nil {
// 	// Send response or just ignore packet
// 	// TODO : DDOS!!??
// 	return
// }
