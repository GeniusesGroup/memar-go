/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"../log"
)

func init() {
	log.Info("-------------------Achaemenid Server-------------------")
	log.Info("Server start at ", time.Now())
}

// DefaultServer use as default server.
var DefaultServer = &defaultServer
var defaultServer Server

// Server represents needed data to serving some functionality such as networks, ...
// to both server and client apps!
type Server struct {
	State           int // States locate in const of this file.
	Manifest        Manifest
	Cryptography    cryptography
	Networks        networks
	StreamProtocols streamProtocols
	Services        services
	Connections     connections
	Nodes           nodes
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
func (s *Server) Init() {
	log.Info("Try to initialize server...")
	if s == nil {
		s = DefaultServer
	}
	s.State = ServerStateStarting
	// Get UserGivenPermission from OS
	s.Assets.init()
	s.Services.init()
	s.Connections.init(s)
	s.Cryptography.init(s)
}

// Start will start the server.
func (s *Server) Start() (err error) {
	log.Info("Try to start server...")

	// Recover from panics if exist to prevent stop server.
	defer s.panicHandler()

	s.State = ServerStateRunning

	// watch for SIGTERM and SIGINT from the operating system, and notify the app on the channel
	var sig = make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM)
	signal.Notify(sig, syscall.SIGINT)
	// Block main goroutine to handle OS signals!
	s.HandleOSSignals(sig)

	return nil
}

func (s *Server) panicHandler() {
	var r = recover()
	if r != nil {
		log.Warn("Panic Exception: ", r)
	}
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
			log.Info("Waiting for server to finish and release proccess, It will take up to 60 seconds")
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

	// Wait to finish above logic, or kill app in --- second!
	// time.Sleep(10 * time.Second)

	// it must change to ServerStateStop(0) otherwise it means app can't close normally
	s.State = ServerStateStop

	log.Info("Server stopped at", time.Now())
	log.SaveToStorage("Achaemenid", repoLocation)
}
