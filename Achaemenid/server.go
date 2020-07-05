/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"os"
	"os/signal"
	"syscall"

	"../log"
)

func init() {
	log.Info("---------------Achaemenid Server---------------")
}

// DefaultServer use as default server.
var DefaultServer = &defaultServer
var defaultServer Server

// Server represents needed data to serving some functionality such as networks, ...
// to both server and client apps!
type Server struct {
	State           int // States locate in const of this file.
	Manifest        Manifest
	Networks        networks
	StreamProtocols streamProtocols
	Cryptography    cryptography
	Services        services
	Connections     connections
	Assets          assets
}

// Server State
const (
	ServerStateStop int = iota
	ServerStateRunning
	ServerStateStopping
	ServerStateStarting // plan to start
)

// Init method use to initialize server object with default data.
func (s *Server) Init() (err error) {
	log.Info("Try to initialize server...")

	if s == nil {
		s = DefaultServer
	}
	s.Assets.init()
	s.Services.init()
	s.Connections.init()

	// Make & Register cryptography data
	err = s.Cryptography.init(s)
	return
}

// Start will start the server.
func (s *Server) Start() (err error) {
	log.Info("Try to start server...")

	s.State = ServerStateStarting

	// Recover from panics if exist.
	// defer panicHandler(s)

	// Get UserGivenPermission from OS

	s.State = ServerStateRunning

	// register s.HandleGP() for income packet handler

	// watch for SIGTERM and SIGINT from the operating system, and notify the app on the channel
	var sig = make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM)
	signal.Notify(sig, syscall.SIGINT)
	// Block main goroutine to handle OS signals!
	s.HandleOSSignals(sig)

	return nil
}

// HandleOSSignals use to handle OS signals! Caller will block until we get an OS signal
func (s *Server) HandleOSSignals(sigChannel chan os.Signal) {
	var sig = <-sigChannel
	log.Warn("caught signal: ", sig)
	switch sig {
	// wait for our os signal to stop the app on the graceful stop channel
	case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL:
		if s.State == ServerStateRunning {
			s.State = ServerStateStopping
			log.Warn("Waiting for server to finish and release proccess, It will take up to 60 seconds")
			s.Shutdown()
			os.Exit(s.State)
		}
	}
}

// Shutdown use to graceful stop server!!
func (s *Server) Shutdown() {
	s.Cryptography.shutdown()
	s.Networks.shutdown()
	s.Assets.shutdown()
	log.SaveToStorage("Achaemenid", repoLocation)

	// Wait to finish above logic, or kill app in --- second!
	// time.Sleep(10 * time.Second)

	// it must change to ServerStateStop(0) otherwise it means app can't close normally
	s.State = ServerStateStop
}

// SendStream use to register a stream to send pool and automatically send to the peer.
func (s *Server) SendStream(st *Stream) (err error) {
	// First Check st.Connection.Status to ability send stream over it

	return nil
}
