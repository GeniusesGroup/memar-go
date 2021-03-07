/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"crypto/sha512"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"syscall"
	"time"

	as "../assets"
	"../convert"
	"../log"
)

// Server is the base object that use by other part of app and platforms!
var Server serverStructure

// Init method use to auto initialize server object with default data.
func init() {
	// Indicate repoLocation
	// TODO::: change to PersiaOS when it ready!
	var ex, err = os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	Server.RepoLocation = filepath.Dir(ex)

	Server.Services.init()

	// Server.Assets.GUI = as.NewFolder("gui")
	Server.Assets.Objects = as.NewFolder("objects")
	Server.Assets.Secret = as.NewFolder("secret")
	Server.Assets.Secret.ReadRepositoryFromFileSystem(Server.RepoLocation+"/secret", true)
}

// serverStructure represents needed data to serving some functionality such as networks, ...
// to both server and client apps!
type serverStructure struct {
	AppID           [32]byte // Hash of domain act as Application ID too
	State           int      // States locate in const of this file.
	RepoLocation    string   // Just fill in supported OS
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

// Init method use to initialize server object with default data in second phase.
func (s *serverStructure) Init() {
	defer s.PanicHandler()

	log.Init("Achaemenid", Server.RepoLocation, 24*60*60)

	log.Info("-----------------------------Achaemenid Server-----------------------------")
	log.Info("Try to initialize Achaemenid Server...")

	Server.State = ServerStateStarting

	// Get UserGivenPermission from OS

	log.Info("App start in", Server.RepoLocation)

	s.Manifest.init()

	s.State = ServerStateStarting

	// Get UserGivenPermission from OS

	s.AppID = sha512.Sum512_256(convert.UnsafeStringToByteSlice(s.Manifest.DomainName))

	Server.Assets.LoadFromStorage()
	Server.Connections.init()
	s.Cryptography.init()
}

// Start will start the server and block caller until server shutdown.
func (s *serverStructure) Start() {
	log.Info("Try to start achaemenid.Server...")

	s.State = ServerStateRunning

	log.Info("Server start successfully, Now listen to any OS signals ...")

	// watch for SIGTERM and SIGINT from the operating system, and notify the app on the channel
	// Infinity loop block main goroutine to handle OS signals!
	for {
		var sig = make(chan os.Signal, 1)
		signal.Notify(sig)
		s.HandleOSSignals(sig)
	}
}

// PanicHandler recover from panics if exist to prevent server stop.
// Call it by defer in any goroutine due to >> https://github.com/golang/go/issues/20161
func (s *serverStructure) PanicHandler() {
	var r = recover()
	if r != nil {
		log.Warn("Panic Exception: ", r)
		log.Warn("Debug Stack: ", string(debug.Stack()))
	}
}

// HandleOSSignals use to handle OS signals! Caller will block until we get an OS signal
// https://en.wikipedia.org/wiki/Signal_(IPC)
func (s *serverStructure) HandleOSSignals(sigChannel chan os.Signal) {
	var sig = <-sigChannel
	switch sig {
	case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL:
		log.Info("Caught signal to stop server")
		if s.State == ServerStateRunning {
			var timer = time.NewTimer(s.Manifest.TechnicalInfo.ShutdownDelay)
			s.State = ServerStateStopping
			log.Info("Waiting for server to finish and release proccess, It will take up to 60 seconds")
			go s.Shutdown()

			select {
			case <-timer.C:
				if s.State == ServerStateStopping {
					log.Info("Server can't finish shutdown and release proccess in", s.Manifest.TechnicalInfo.ShutdownDelay.String())
				} else {
					log.Info("Server stopped successfully")
				}
			}

			log.SaveToStorage()
			os.Exit(s.State)
		}
	case syscall.Signal(0x17): // syscall.SIGURG:
		log.Warn("Caught urgened signal: ", sig)
	case syscall.Signal(10): // syscall.SIGUSR1
		log.Info("Caught signal to reload server assets")
		s.Assets.LoadFromStorage()
	// case syscall.Signal(12): // syscall.SIGUSR2
	default:
		log.Warn("Caught un-managed signal: ", sig)
	}
}

// Shutdown use to graceful stop server!!
func (s *serverStructure) Shutdown() {
	s.State = ServerStateStopping

	s.Cryptography.shutdown()
	s.Networks.shutdown()
	s.Assets.shutdown()
	s.Connections.shutdown()

	// Wait to finish above logic, or kill app in --- second!
	// time.Sleep(10 * time.Second)

	// it must change to ServerStateStop(0) otherwise it means app can't close normally
	s.State = ServerStateStop
}
