/* For license and copyright information please see the LEGAL file in the code repository */

package achaemenid

import (
	"strconv"

	"memar/errors"
	"memar/event"
	"memar/log"
	"memar/net"
	"memar/protocol"
	"memar/service"
)

// App is the base object that use by other part of app and platforms.
// It is implement protocol.Application interface
var App app

// app represents application (server or client) requirements data to serving some functionality such as networks, ...
type app struct {
	softwareStatus protocol.SoftwareStatus
	status

	manifest Manifest

	Cryptography cryptography

	service.SS
	errors.Errors
	event.EventTarget
	net.PacketListener
}

func (app *app) Engine() protocol.ApplicationEngine      { return nil }
func (app *app) Manifest() protocol.ApplicationManifest  { return &app.manifest }
func (app *app) SoftwareStatus() protocol.SoftwareStatus { return app.softwareStatus }

// Init method use to initialize app object with default data in second phase.
//
//memar:impl memar/protocol.ObjectLifeCycle
func (app *app) Init() (err protocol.Error) {
	defer log.PanicHandler()

	log.Info(&domain, `-----------------------------Achaemenid Application-----------------------------
Try to initialize application...`)

	app.changeState(protocol.ApplicationStatus_Starting)

	app.SS.Init()
	app.Errors.Init()
	err = app.EventTarget.Init()
	if err != nil {
		return
	}

	// Get UserGivenPermission from OS

	app.manifest.init()
	err = app.Cryptography.init()
	err = app.PacketListener.Init()
	return
}

// Reinit or Reload use to reload application
func (app *app) Reinit() (err protocol.Error) {
	return
}

// Deinit use to graceful stop app
func (app *app) Deinit() (err protocol.Error) {
	app.changeState(protocol.ApplicationStatus_Stopping)
	log.Info(&domain, "Waiting for app to finish and release process, It will take up to 60 seconds")

	app.changeState(protocol.ApplicationStatus_Stopping)

	err = app.Cryptography.Deinit()
	if err != nil {
		return
	}

	// write files to storage device if any change made

	// Wait to finish above logic, or kill app in --- second
	// time.Sleep(10 * time.Second)

	err = app.PacketListener.Deinit()
	return
}

//memar:impl memar/protocol.OS_Signal_Listener
func (app *app) OsSignalHandler(signal protocol.OS_Signal) {
	switch signal {
	case protocol.OS_Signal_Interrupt, protocol.OS_Signal_Quit, protocol.OS_Signal_Terminated, protocol.OS_Signal_Killed:
		log.Info(&domain, "Caught signal to stop app")
		if app.Status() == protocol.ApplicationStatus_Running {
			go app.Deinit()
		}
	case protocol.OS_Signal_Urgent:
		log.Warn(&domain, "Caught urgent I/O condition signal")
	case protocol.OS_Signal_Hangup:
		log.Info(&domain, "Caught signal to reload app")
		app.Reinit()
	case protocol.OS_Signal_Upgrade:
		log.Info(&domain, "Caught signal to update, upgrade and reload application")
		app.Upgrade()
	default:
		var msg = "Caught un-managed signal: " + strconv.Itoa(int(signal))
		log.Warn(&domain, msg)
	}
}

// Upgrade use to update, upgrade and reload application
func (app *app) Upgrade() {
	// TODO::: update, upgrade application
	// https://github.com/cloudflare/tableflip

	app.Reinit()
}
