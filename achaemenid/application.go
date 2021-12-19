/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	goos "os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"../connection"
	er "../error"
	"../ganjine"
	"../log"
	"../node"
	"../protocol"
	"../service"
)

// App is the base object that use by other part of app and platforms!
// It is implement protocol.Application interface
var App appStructure

// appStructure represents application (server or client) requirements data to serving some functionality such as networks, ...
type appStructure struct {
	softwareStatus    protocol.SoftwareStatus
	state             protocol.ApplicationState
	stateListeners    []chan protocol.ApplicationState
	stateChangeLocker sync.Mutex

	Cryptography cryptography

	ganjine.Cluster
	log.Logger
	node.Nodes
	service.Services
	er.Errors
	connection.Connections

	netAppMux
}

func (app *appStructure) SoftwareStatus() protocol.SoftwareStatus { return app.softwareStatus }
func (app *appStructure) Status() protocol.ApplicationState       { return app.state }
func (app *appStructure) NotifyState(notifyBy chan protocol.ApplicationState) {
	app.stateListeners = append(app.stateListeners, notifyBy)
}

// Init method use to initialize app object with default data in second phase.
func (app *appStructure) Init() {
	defer app.PanicHandler()

	protocol.App.Log(log.InfoEvent(domainEnglish, "-----------------------------Achaemenid Application-----------------------------"))
	protocol.App.Log(log.InfoEvent(domainEnglish, "Try to initialize application..."))

	app.changeState(protocol.ApplicationStateStarting)

	if app.Manifest.DataNodes.TotalZones < 3 {
		protocol.App.Log(protocol.Log_Warning, "ReplicationNumber set below 3! Loose write ability until two replication available again on replication failure!")
	}

	app.Services.Init()
	app.Errors.Init()
	app.Connections.Init()

	// Get UserGivenPermission from OS

	app.Manifest.init()
	app.Cryptography.init()

	registerServices()
}

// Start will start the app and block caller until app shutdown.
func (app *appStructure) Start() {
	protocol.App.Log(log.InfoEvent(domainEnglish, "Try to start application ..."))

	app.changeState(protocol.ApplicationStateRunning)

	protocol.App.Log(log.InfoEvent(domainEnglish, "Application start successfully, Now listen to any OS signals ..."))

	// Block main goroutine to handle OS signals.
	var sig = make(chan goos.Signal, 1024)
	signal.Notify(sig)
	app.HandleOSSignals(sig)
}

// HandleOSSignals use to handle OS signals! Caller will block until app terminate or exit.
// https://en.wikipedia.org/wiki/Signal_(IPC)
func (app *appStructure) HandleOSSignals(sigChannel chan goos.Signal) {
	for {
		var sig = <-sigChannel
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL:
			protocol.App.Log(log.InfoEvent(domainEnglish, "Caught signal to stop app"))
			if app.State() == protocol.ApplicationStateRunning {
				go app.Shutdown()
			}
		case syscall.Signal(0x17): // syscall.SIGURG:
			protocol.App.Log(protocol.Log_Warning, "Caught urgened signal: "+sig)
		case syscall.Signal(10): // syscall.SIGUSR1
			protocol.App.Log(log.InfoEvent(domainEnglish, "Caught signal to reload app"))
			app.Reload()
		// case syscall.Signal(12): // syscall.SIGUSR2
		default:
			protocol.App.Log(protocol.Log_Warning, "Caught un-managed signal: "+sig)
		}
	}
}

// Reload use to reload app
func (app *appStructure) Reload() {}

// Shutdown use to graceful stop app
func (app *appStructure) Shutdown() {
	app.changeState(protocol.ApplicationStateStopping)
	protocol.App.Log(log.InfoEvent(domainEnglish, "Waiting for app to finish and release proccess, It will take up to 60 seconds"))

	app.changeState(protocol.ApplicationStateStopping)

	app.Cryptography.shutdown()
	app.Connections.Shutdown()
	app.Nodes.Shutdown()

	// write files to storage device if any change made

	// Wait to finish above logic, or kill app in --- second
	// time.Sleep(10 * time.Second)
	var timer = time.NewTimer(app.Manifest.TechnicalInfo.ShutdownDelay)
	select {
	case <-timer.C:
		if app.Status() == protocol.ApplicationStateStopping {
			protocol.App.Log(log.InfoEvent(domainEnglish, "Application can't finish shutdown and release proccess in"+app.Manifest.TechnicalInfo.ShutdownDelay.String()))
			goos.Exit(1)
		} else {
			protocol.App.Log(log.InfoEvent(domainEnglish, "Application stopped successfully"))
			goos.Exit(0)
		}
	}

}

// Shutdown use to graceful stop app
func (app *appStructure) changeState(state protocol.ApplicationState) {
	app.stateChangeLocker.Lock()
	app.state = state
	for _, listener := range app.stateListeners {
		listener <- state
	}
	app.stateChangeLocker.Unlock()
}
