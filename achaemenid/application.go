/* For license and copyright information please see the LEGAL file in the code repository */

package achaemenid

import (
	"strconv"
	"sync"

	er "libgo/error"
	"libgo/event"
	"libgo/log"
	"libgo/protocol"
	"libgo/service"
)

// App is the base object that use by other part of app and platforms.
// It is implement protocol.Application interface
var App app

// app represents application (server or client) requirements data to serving some functionality such as networks, ...
type app struct {
	softwareStatus    protocol.SoftwareStatus
	state             protocol.ApplicationState
	stateListeners    []chan protocol.ApplicationState
	stateChangeLocker sync.Mutex

	Manifest Manifest

	Cryptography cryptography

	log.Logger
	service.SS
	er.Errors
	event.EventTarget
	netAppMux
}

func (app *app) Engine() protocol.ApplicationEngine      { return nil }
func (app *app) SoftwareStatus() protocol.SoftwareStatus { return app.softwareStatus }
func (app *app) Status() protocol.ApplicationState       { return app.state }
func (app *app) NotifyState(notifyBy chan protocol.ApplicationState) {
	app.stateListeners = append(app.stateListeners, notifyBy)
}

// Init method use to initialize app object with default data in second phase.
//
//libgo:impl libgo/protocol.ObjectLifeCycle
func (app *app) Init() (err protocol.Error) {
	defer app.PanicHandler()

	protocol.App.Log(log.InfoEvent(&DefaultEvent_MediaType, `-----------------------------Achaemenid Application-----------------------------
Try to initialize application...`))

	app.changeState(protocol.ApplicationState_Starting)

	app.SS.Init()
	app.Errors.Init()
	err = app.EventTarget.Init()
	if err != nil {
		return
	}

	// Get UserGivenPermission from OS

	app.Manifest.init()
	app.Cryptography.init()
	return
}

// Reinit or Reload use to reload application
func (app *app) Reinit() {}

// Deinit use to graceful stop app
func (app *app) Deinit() (err protocol.Error) {
	app.changeState(protocol.ApplicationState_Stopping)
	protocol.App.Log(log.InfoEvent(&DefaultEvent_MediaType, "Waiting for app to finish and release process, It will take up to 60 seconds"))

	app.changeState(protocol.ApplicationState_Stopping)

	err = app.Cryptography.Deinit()
	if err != nil {
		return
	}

	// write files to storage device if any change made

	// Wait to finish above logic, or kill app in --- second
	// time.Sleep(10 * time.Second)
	return
}

//libgo:impl libgo/protocol.OS_Signal_Listener
func (app *app) OsSignalHandler(signal protocol.OS_Signal) {
	switch signal {
	case protocol.OS_Signal_Interrupt, protocol.OS_Signal_Quit, protocol.OS_Signal_Terminated, protocol.OS_Signal_Killed:
		protocol.App.Log(log.InfoEvent(&DefaultEvent_MediaType, "Caught signal to stop app"))
		if app.Status() == protocol.ApplicationState_Running {
			go app.Deinit()
		}
	case protocol.OS_Signal_Urgent:
		protocol.App.Log(log.WarnEvent(&DefaultEvent_MediaType, "Caught urgent I/O condition signal"))
	case protocol.OS_Signal_Hangup:
		protocol.App.Log(log.InfoEvent(&DefaultEvent_MediaType, "Caught signal to reload app"))
		app.Reinit()
	case protocol.OS_Signal_Upgrade:
		protocol.App.Log(log.InfoEvent(&DefaultEvent_MediaType, "Caught signal to update, upgrade and reload application"))
		app.Upgrade()
	default:
		var msg = "Caught un-managed signal: " + strconv.Itoa(int(signal))
		protocol.App.Log(log.WarnEvent(&DefaultEvent_MediaType, msg))
	}
}

// Upgrade use to update, upgrade and reload application
func (app *app) Upgrade() {
	// TODO::: update, upgrade application
	// https://github.com/cloudflare/tableflip

	app.Reinit()
}

func (app *app) changeState(state protocol.ApplicationState) {
	app.stateChangeLocker.Lock()
	defer app.stateChangeLocker.Unlock()
	app.state = state
	for _, listener := range app.stateListeners {
		// TODO::: can be blocking if listener hasn't enough capacity buffer??
		listener <- state
	}
}
