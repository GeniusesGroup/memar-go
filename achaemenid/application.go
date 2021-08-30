/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"../connections"
	er "../error"
	"../log"
	"../os"
	"../protocol"
	"../service"
)

// App is the base object that use by other part of app and platforms!
var App appStructure

// appStructure represents needed data to serving some functionality such as networks, ...
// to both server and client apps!
type appStructure struct {
	State        protocol.ApplicationState
	StateChannel chan protocol.ApplicationState
	Manifest     Manifest
	Cryptography cryptography

	os.OS
	netAppMux
	service.Services
	connections.Connections
	protocol.Errors

	// fileDirectory   file.Directory
	// objectDirectory object.Directory

	Nodes nodes
}

// Init method use to initialize app object with default data in second phase.
func (app *appStructure) Init() {
	defer app.PanicHandler()

	log.Info("-----------------------------Achaemenid Application-----------------------------")
	log.Info("Try to initialize Achaemenid Application...")

	app.State = protocol.ApplicationStateStarting

	os.OS.Init()

	App.Services.Init()

	log.Init("Achaemenid", 24*60*60)

	// Get UserGivenPermission from OS

	app.Manifest.init()

	// Get UserGivenPermission from OS

	app.Connections.init()
	app.Cryptography.init()
	app.Errors = &er.Errors
}

// Start will start the app and block caller until app shutdown.
func (app *appStructure) Start() {
	log.Info("Try to start achaemenid.App...")

	app.State = protocol.ApplicationStateRunning

	log.Info("Application start successfully, Now listen to any OS signals ...")

	// watch for SIGTERM and SIGINT from the operating system, and notify the app on the channel
	// Infinity loop block main goroutine to handle OS signals!
	for {
		var sig = make(chan os.Signal, 1)
		signal.Notify(sig)
		app.HandleOSSignals(sig)
	}
}

func (app *appStructure) OS() protocol.OS       { return &app.OS }
func (app *appStructure) Manifest() protocol.OS { return &app.Manifest }

// PanicHandler recover from panics if exist to prevent app stop.
// Call it by defer in any goroutine due to >> https://github.com/golang/go/issues/20161
func (app *appStructure) PanicHandler() {
	var r = recover()
	if r != nil {
		log.Warn("Panic Exception: ", r)
		log.Warn("Debug Stack: ", string(debug.Stack()))
	}
}

// HandleOSSignals use to handle OS signals! Caller will block until we get an OS signal
// https://en.wikipedia.org/wiki/Signal_(IPC)
func (app *appStructure) HandleOSSignals(sigChannel chan os.Signal) {
	var sig = <-sigChannel
	switch sig {
	case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL:
		log.Info("Caught signal to stop app")
		if app.State == protocol.ApplicationStateRunning {
			var timer = time.NewTimer(app.Manifest.TechnicalInfo.ShutdownDelay)
			app.State = protocol.ApplicationStateStopping
			log.Info("Waiting for app to finish and release proccess, It will take up to 60 seconds")
			go app.Shutdown()

			select {
			case <-timer.C:
				if app.State == protocol.ApplicationStateStopping {
					log.Info("Application can't finish shutdown and release proccess in", app.Manifest.TechnicalInfo.ShutdownDelay.String())
				} else {
					log.Info("Application stopped successfully")
				}
			}

			log.SaveToStorage()
			os.Exit(app.State)
		}
	case syscall.Signal(0x17): // syscall.SIGURG:
		log.Warn("Caught urgened signal: ", sig)
	case syscall.Signal(10): // syscall.SIGUSR1
		log.Info("Caught signal to reload app file")
		app.Reload()
	// case syscall.Signal(12): // syscall.SIGUSR2
	default:
		log.Warn("Caught un-managed signal: ", sig)
	}
}

// Reload use to reload app
func (app *appStructure) Reload() {}

// Shutdown use to graceful stop app
func (app *appStructure) Shutdown() {
	app.State = protocol.ApplicationStateStopping
	app.StateChannel = protocol.ApplicationStateStopping

	app.Cryptography.shutdown()
	app.Connections.Shutdown()

	// write files to storage device if any change made

	// Wait to finish above logic, or kill app in --- second!
	// time.Sleep(10 * time.Second)

	// it must change to protocol.ApplicationStateStop(0) otherwise it means app can't close normally
	app.State = protocol.ApplicationStateStop
}
