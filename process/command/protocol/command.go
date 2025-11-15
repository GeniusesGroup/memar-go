/* For license and copyright information please see the LEGAL file in the code repository */

package command_p

import (
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	mediatype_p "memar/identifier/mediatype/protocol"
	request_p "memar/process/request/protocol"
	response_p "memar/process/response/protocol"
)

// Command is the interface that must implement by any struct to be a command service
// It is old user interface, In `uni-kernel` and `Graphic UI` it is useless because there is no `terminal` and `shell`.
// TODO::: Deprecate in favor of service??
// 
// https://en.wikipedia.org/wiki/Command-line_interface
type Command interface {
	// Init(parent Command, subCommands ...Command)

	// RunnableCommand reports whether the command can be run; otherwise it is a documentation pseudo-command
	RunnableCommand() bool

	// ParentCommand return the parent command for this command.
	// It can be nill for the root command.
	ParentCommand() Command
	// SubCommand return a sub command by its name or alias that must use intelligent suggestion
	SubCommand(name string) Command
	// Commands lists the available commands and help topics.
	// The order here is the order in which they are printed by 'go help'.
	// Note that subcommands are in general best avoided.
	SubCommands() []Command

	Handler

	datatype_p.DataType
	mediatype_p.MediaType

	request_p.Request
	response_p.Response
}

// Handler introduce CLI (command-line interface) service handler.
type Handler interface {
	// ServeCLA or serve by command-line arguments might block the caller
	// Arguments list not include the command name.
	ServeCLA(args Arguments) (err error_p.Error)

	// read and write to e.g. os.Stdin, os.Stdout, and os.Stderr files
	// ServeCLI() (err error_p.Error)
}
