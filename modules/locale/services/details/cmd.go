/* For license and copyright information please see the LEGAL file in the code repository */

package locale

import (
	cmd "memar/process/command"
	"memar/process/service"
	// "memar/modules/app/services"
)

var Command command

type command struct {
	service.Service
	cmd.Command
}

// Init use to initialize all services and register them in the given parent.
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (c *command) Init(parent command_p.Command) (err error_p.Error) {
	// c.Command.Init("locale", "memar locale <command> [arguments]", parent, false)

	// services.Init.Init(c)

	// c.Command.SetSubCommands(&services.Init)

	return
}
